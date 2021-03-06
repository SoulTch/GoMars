package core

import (
	"fmt"
)

// game configurations
const (
	maxPlayer = 5
)

// GameObject interface
type GameObject interface {
	refresh()
}

// Game struct.
type Game struct {
	numPlayer    int
	players      []*player
	board        board
	gateway      *Gateway
	eventHandler [][]handler

	tags []int

	enchants map[string]int
}

// GameResult struct.
type GameResult struct {
}

// CreateGame creates game and initiates.
func CreateGame(playerCount int) (*Game, error) {
	if playerCount > maxPlayer {
		return nil, fmt.Errorf("player exceeded maximum player number : %d", playerCount)
	}

	game := new(Game)
	game.players = make([]*player, playerCount, playerCount)
	game.gateway = CreateGateway(game)
	game.eventHandler = createEventHandler()
	game.enchants = make(map[string]int)
	game.tags = make([]int, tagsize)

	for i := 0; i < playerCount; i++ {
		game.players[i] = createPlayer(game, i)
	}

	return game, nil
}

func (game *Game) startGame() GameResult {

	return GameResult{}
}

func (game *Game) refresh() {
	for _, i := range game.players {
		i.refresh()
	}
}

func (game *Game) getEnchant() map[string]int {
	return game.enchants
}

func (game *Game) invoke(e event) {
	oldHandlers := &game.eventHandler[e.getEventID()]
	newHandlers := make([]handler, 0, len(*oldHandlers))

	for _, i := range *oldHandlers {
		if !i.handle(game, e) {
			newHandlers = append(newHandlers, i)
		}
	}

	oldHandlers = &newHandlers
}
