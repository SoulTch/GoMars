package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/SoulTch/GoMars/core/protocol"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
)

type client struct {
	hub  *hub
	pID  int
	conn *websocket.Conn
}

func buildClient(hub *hub, pID int, conn *websocket.Conn) *client {
	c := new(client)
	c.hub = hub
	c.pID = pID
	c.conn = conn

	return c
}

func authConnection(key string, conn *websocket.Conn) error {
	// authentication
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Print("authentication failed. reading faild.")
		return err
	}

	if xkey := string(bytes.TrimSpace(message)); key != xkey {
		return fmt.Errorf("authentication failed. expected %s, got %s", key, xkey)
	}

	return nil
}

func (c *client) run() {
	go func() {
		err := authConnection(c.hub.pID[c.pID], c.conn)
		if err != nil {
			log.Print(err.Error())
			c.conn.Close()
			return
		}

		// server -> client
		go func() {
			ticker := time.NewTicker(pingPeriod)
			defer func() {
				ticker.Stop()
				c.conn.Close()
			}()

			for {
				select {
				case message, ok := <-c.hub.gateway.ActionChan[c.pID]:
					if !ok {
						c.conn.WriteMessage(websocket.CloseMessage, []byte{})
						return
					}

					c.conn.SetWriteDeadline(time.Now().Add(writeWait))
					err := c.conn.WriteJSON(message)
					if err != nil {
						return
					}

				case message, ok := <-c.hub.gateway.NoticeChan[c.pID]:
					if !ok {
						c.conn.WriteMessage(websocket.CloseMessage, []byte{})
						return
					}
					c.conn.SetWriteDeadline(time.Now().Add(writeWait))
					err := c.conn.WriteJSON(message)
					if err != nil {
						return
					}

				case <-ticker.C:
					c.conn.SetWriteDeadline(time.Now().Add(writeWait))
					if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				}
			}
		}()

		// client -> server
		go func() {
			defer func() {
				c.conn.Close()
			}()
			c.conn.SetReadLimit(maxMessageSize)
			c.conn.SetReadDeadline(time.Now().Add(pongWait))
			c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

			for {
				message := protocol.Response{}
				err := c.conn.ReadJSON(&message)
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("error: %v", err)
					}
					break
				}

				message.Player = c.pID

				switch message.Action {
				case "get":
					c.recv <- message
				case "action":
					c.resp <- message
				}
			}
		}()
	}()
}
