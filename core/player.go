package core

type player struct {
	id   int
	game *Game

	eventHandler [][]playerHandler
	tags         []int

	// resources
	terraforming int
	resources    []int

	// productions
	production []int

	// trades
	mcPerResource []int
	cardCost      int

	// reduces
	globalReduce int

	// projects
	hand    map[int]*project
	board   []*project
	actions []*action
	basics  []*basicProject

	enchants map[string]int
}

type belongsToPlayer interface {
	getPlayer() *player
}

func createPlayer(game *Game, id int) *player {
	player := new(player)
	player.game = game
	player.id = id
	player.eventHandler = createPlayerEventHandler()
	player.tags = make([]int, tagsize)
	player.hand = make(map[int]*project)
	player.actions = make([]*action, 0)
	player.basics = make([]*basicProject, 0)
	player.enchants = make(map[string]int)
	player.resources = make([]int, resSize)
	player.production = make([]int, resSize)

	return player
}

func (player *player) refresh() {

}

func (player *player) getEnchant() map[string]int {
	return player.enchants
}

func (player *player) invoke(e event) {
	player.game.invoke(e)

	oldHandlers := &player.eventHandler[e.getEventID()]
	newHandlers := make([]playerHandler, 0, len(*oldHandlers))

	for _, i := range *oldHandlers {
		if !i.handle(player.game, player, e) {
			newHandlers = append(newHandlers, i)
		}
	}

	oldHandlers = &newHandlers
}
