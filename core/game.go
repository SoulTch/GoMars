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
	players      []*Player
	gateway      *Gateway
	eventHandler [][]Handler

	enchants map[string]int
}

// GameResult struct.
type GameResult struct {
}

// CreateGame creates game and initiates.
func CreateGame(player int) (*Game, error) {
	if player > maxPlayer {
		return nil, fmt.Errorf("player exceeded maximum player number : %d", player)
	}

	game := new(Game)
	game.players = make([]*Player, player, player)
	game.gateway = CreateGateway(game)
	game.eventHandler = CreateEventHandler()
	game.enchants = make(map[string]int)

	for i := 0; i < player; i++ {
		game.players[i] = CreatePlayer(game, i)
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

func (game *Game) invokeEvent(e Event) {
	oldHandlers := &game.eventHandler[int(e.etype)]
	newHandlers := make([]Handler, 0, len(*oldHandlers))

	for _, i := range *oldHandlers {
		if !i.handle(game, e) {
			newHandlers = append(newHandlers, i)
		}
	}

	oldHandlers = &newHandlers
}

func (game *Game) invokeSimpleEvent(e EventType) {
	game.invokeEvent(Event{etype: e})
}
