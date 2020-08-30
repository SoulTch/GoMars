package core

type event interface {
	getEventID() int
}

type handler interface {
	handle(*Game, event) bool
}

type playerHandler interface {
	handle(*Game, *player, event) bool
}

func createEventHandler() [][]handler {
	h := make([][]handler, eventSize)

	for i := 0; i < eventSize; i++ {
		h[i] = make([]handler, 0)
	}

	return h
}

func createPlayerEventHandler() [][]playerHandler {
	h := make([][]playerHandler, eventSize)

	for i := 0; i < eventSize; i++ {
		h[i] = make([]playerHandler, 0)
	}

	return h
}

// event definitions

type startGameEvent struct{}

func (x *startGameEvent) getEventID() int { return 0 }

type endGameEvent struct{}

func (x *endGameEvent) getEventID() int { return 1 }

type generationBeginEvent struct{}

func (x *generationBeginEvent) getEventID() int { return 2 }

type generationEndEvent struct{}

func (x *generationEndEvent) getEventID() int { return 3 }

type addResourcesEvent struct {
	player *player
	rtype  resType
	amount int
}

func (x *addResourcesEvent) getEventID() int { return 4 }

type incProductionEvent struct {
	player *player
	rtype  resType
	amount int
}

func (x *incProductionEvent) getEventID() int { return 5 }

type decProductionEvent struct {
	player *player
	rtype  resType
	amount int
}

func (x *decProductionEvent) getEventID() int { return 6 }

type addCardEvent struct {
	player *player
	amount int
}

func (x *addCardEvent) getEventID() int { return 7 }

type playProjectEvent struct {
	player  *player
	project *project
}

func (x *playProjectEvent) getEventID() int { return 8 }

var eventSize = 9
