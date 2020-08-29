package core

type player struct {
	id   int
	game *Game

	eventHandler [][]playerHandler
	tags         tags

	// resources
	terraforming int
	megacredit   int
	iron         int
	titanium     int
	energy       int
	heat         int
	plant        int

	// productions
	pMegacredit int
	pIron       int
	pTitanium   int
	pEnerty     int
	pHeat       int
	pPlant      int

	// trades
	mcPerIron     int
	mcPerTitanium int
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
	player.hand = make(map[int]*project)
	player.actions = make([]*action, 0)
	player.basics = make([]*basicProject, 0)
	player.enchants = make(map[string]int)

	return player
}

func (player *player) refresh() {

}

func (player *player) getEnchant() map[string]int {
	return player.enchants
}

func (player *player) invokeEvent(e event) {
	player.game.invokeEvent(e)

	oldHandlers := &player.eventHandler[int(e.etype)]
	newHandlers := make([]playerHandler, 0, len(*oldHandlers))

	for _, i := range *oldHandlers {
		if !i.handle(player.game, player, e) {
			newHandlers = append(newHandlers, i)
		}
	}

	oldHandlers = &newHandlers
}

func (player *player) invokeSimpleEvent(e eventType) {
	player.invokeEvent(event{etype: e, player: player})
}
