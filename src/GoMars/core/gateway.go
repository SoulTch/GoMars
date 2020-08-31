package core

import (
	"github.com/SoulTch/GoMars/core/protocol"
)

// Gateway .
type Gateway struct {
	Game       *Game
	ActionChan []chan protocol.Message
	ResChan    []chan protocol.Response

	NoticeChan []chan protocol.Message
	GetChan    []chan protocol.Response
}

// CreateGateway .
func CreateGateway(game *Game) *Gateway {
	numPlayer := game.numPlayer
	makeMessageChan := func() []chan protocol.Message {
		r := make([]chan protocol.Message, numPlayer)
		for i := 0; i < numPlayer; i++ {
			r[i] = make(chan protocol.Message, 10)
		}

		return r
	}

	makeReponseChan := func() []chan protocol.Response {
		r := make([]chan protocol.Response, numPlayer)
		for i := 0; i < numPlayer; i++ {
			r[i] = make(chan protocol.Response, 10)
		}

		return r
	}

	g := new(Gateway)
	g.Game = game
	g.ActionChan = makeMessageChan()
	g.ResChan = makeReponseChan()

	g.NoticeChan = makeMessageChan()
	g.GetChan = makeReponseChan()

	return g
}

func (g *Gateway) action(player int, action string, req interface{}, check func(map[string]int) bool) map[string]int {
	for {
		// send
		g.ActionChan[player] <- protocol.BuildActionMessage(player, action, req)

		// recv
		res := <-g.ResChan[player]

		// response
		if check(res.Data) {
			g.ActionChan[player] <- protocol.BuildResultMessage(player, true)
			return res.Data
		}
		g.ActionChan[player] <- protocol.BuildResultMessage(player, false)
	}
}

func (g *Gateway) notice(player int, action string, noti interface{}) {
	if player >= 0 {
		g.ActionChan[player] <- protocol.BuildNoticeMessage(player, action, noti)
	}

	for i := 0; i < g.Game.numPlayer; i++ {
		g.ActionChan[i] <- protocol.BuildNoticeMessage(player, action, noti)
	}
}
