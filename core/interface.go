package core

// You should change game state only by calling these methods.
// DO NOT CHANGE ANY VALUE DIRECTLY

func (g *Game) gameBegin() {
	g.invokeEvent(Event{etype: startGameEvent})
}

func (g *Game) gameEnd() {
	g.invokeEvent(Event{etype: endGameEvent})
}
