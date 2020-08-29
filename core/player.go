package core

// Player struct.
type Player struct {
	id   int
	game *Game

	eventHandler [][]PlayerHandler

	enchants map[string]int
}

// CreatePlayer creates player and initiates.
func CreatePlayer(game *Game, id int) *Player {
	player := new(Player)
	player.game = game
	player.id = id
	player.eventHandler = CreatePlayerEventHandler()
	player.enchants = make(map[string]int)

	return player
}

func (player *Player) refresh() {

}

func (player *Player) getEnchant() map[string]int {
	return player.enchants
}

func (player *Player) invokeEvent(e Event) {
	player.game.invokeEvent(e)

	oldHandlers := &player.eventHandler[int(e.etype)]
	newHandlers := make([]PlayerHandler, 0, len(*oldHandlers))

	for _, i := range *oldHandlers {
		if !i.handle(player.game, player, e) {
			newHandlers = append(newHandlers, i)
		}
	}

	oldHandlers = &newHandlers
}

func (player *Player) invokeSimpleEvent(e EventType) {
	player.invokeEvent(Event{etype: e, player: player})
}
