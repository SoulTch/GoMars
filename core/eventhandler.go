package core

type eventType int

const (
	startGameEvent = eventType(iota)
	endGameEvent
	esize
)

var eventSize = int(esize)

type event struct {
	etype  eventType
	player *player
	detail map[string]int
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
