package core

// You should change game state only by calling these methods.
// DO NOT CHANGE ANY VALUE DIRECTLY

// game flow
func (g *Game) gameBegin() {
	g.invoke(&startGameEvent{})
}

func (g *Game) gameEnd() {
	g.invoke(&endGameEvent{})
}

func (g *Game) generationBegin() {
	g.invoke(&generationBeginEvent{})
}

func (g *Game) generationEnd() {
	g.invoke(&generationEndEvent{})
}

// actions
func (p *player) addResources(t resType, amount int) {
	p.resources[int(t)] += amount
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
