package main

import (
	"math/rand"

	"github.com/SoulTch/GoMars/core"
)

var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type hub struct {
	game    *core.Game
	gateway *core.Gateway
	pID     []string

	clients []*client
}

func buildAndStartHub(numPlayer int) *hub {
	h := new(hub)
	h.game, _ = core.CreateGame(numPlayer)
	h.gateway = core.CreateGateway(h.game)

	h.pID = make([]string, numPlayer)
	h.clients = make([]*client, numPlayer)

	for i := 0; i < numPlayer; i++ {
		h.pID[i] = randSeq(32)
	}
}
