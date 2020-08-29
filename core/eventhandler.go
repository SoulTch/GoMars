package core

// EventType .
type EventType int

const (
	startGameEvent = EventType(iota)
	endGameEvent
	esize
)

var eventSize = int(esize)

// Event .
type Event struct {
	etype  EventType
	player *Player
	detail map[string]int
}

// Handler struct
type Handler interface {
	handle(*Game, Event) bool
}

// PlayerHandler .
type PlayerHandler interface {
	handle(*Game, *Player, Event) bool
}

// CreateEventHandler struct.
func CreateEventHandler() [][]Handler {
	h := make([][]Handler, eventSize)

	for i := 0; i < eventSize; i++ {
		h[i] = make([]Handler, 0)
	}

	return h
}

// CreatePlayerEventHandler .
func CreatePlayerEventHandler() [][]PlayerHandler {
	h := make([][]PlayerHandler, eventSize)

	for i := 0; i < eventSize; i++ {
		h[i] = make([]PlayerHandler, 0)
	}

	return h
}
