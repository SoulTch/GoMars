package core

// You should change game state only by calling these methods.
// DO NOT CHANGE ANY VALUE DIRECTLY

const (
	broadcast = -1
)

// game flow
func (g *Game) gameBegin() {
	g.gateway.notice(broadcast, "game begin", nil)
	g.invoke(&startGameEvent{})
}

func (g *Game) gameEnd() {
	g.gateway.notice(broadcast, "game end", nil)
	g.invoke(&endGameEvent{})
}

func (g *Game) generationBegin(gen int) {
	g.gateway.notice(broadcast, "generation begin", gen)
	g.invoke(&generationBeginEvent{})
}

func (g *Game) generationEnd(gen int) {
	g.gateway.notice(broadcast, "generation end", gen)
	g.invoke(&generationEndEvent{})
}

func (p *player) addResources(t resType, amount int) {
	g := p.game

	p.resources[int(t)] += amount
	g.gateway.notice(broadcast, "add resources", map[string]interface{}{
		"target":   p.id,
		"res_type": string(t),
		"amount":   amount})
	p.invoke(&addResourcesEvent{rtype: t, amount: amount, player: p})
}

func (p *player) payResource(t resType, amount int, methods payMethod) {
	// TODO
}

func (p *player) incProduction(t resType, amount int) {
	p.production[int(t)] += amount
	p.invoke(&incProductionEvent{rtype: t, amount: amount, player: p})
}

func (p *player) decProduction(t resType, amount int) {
	p.production[int(t)] -= amount
	p.invoke(&decProductionEvent{rtype: t, amount: amount, player: p})
}

func (p *player) addCard(cards []*project) {
	for _, c := range cards {
		p.hand[c.id] = c
	}
	p.invoke(&addCardEvent{amount: len(cards), player: p})
}

func (p *player) playProject(project *project) {
	p.invoke(&playProjectEvent{player: p, project: project})
}
