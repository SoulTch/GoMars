package core

//ActionRequest .
type ActionRequest interface {
}

//ActionResponse .
type ActionResponse interface {
}

// AskRequest interface
type AskRequest interface {
}

// AskResponse interface
type AskResponse interface {
}

// Notice .
type Notice interface {
}

// GetRequest .
type GetRequest interface {
}

// ActionRequestJob .
type ActionRequestJob struct {
	req ActionRequest
	res chan ActionResponse
}

// AskRequestJob .
type AskRequestJob struct {
	req AskRequest
	res chan AskResponse
}

// Gateway .
type Gateway struct {
	game       *Game
	actionChan chan ActionRequestJob
	askChan    chan AskRequestJob
	getChan    chan GetRequest
	noticeChan chan Notice
}

// CreateGateway .
func CreateGateway(game *Game) *Gateway {
	gateway := new(Gateway)
	gateway.game = game
	gateway.actionChan = make(chan ActionRequestJob)
	gateway.askChan = make(chan AskRequestJob)
	gateway.getChan = make(chan GetRequest)
	gateway.noticeChan = make(chan Notice)

	return gateway
}

// Action .
func (g *Gateway) Action(req ActionRequest) ActionResponse {
	res := make(chan ActionResponse, 1)
	reqJob := ActionRequestJob{req, res}
	g.actionChan <- reqJob
	return <-res
}

// Ask .
func (g *Gateway) Ask(req AskRequest) AskResponse {
	res := make(chan AskResponse, 1)
	reqJob := AskRequestJob{req, res}
	g.getChan <- reqJob
	return <-res
}

// Notice .
func (g *Gateway) Notice(noti Notice) {
	g.noticeChan <- noti
}

// Get .
func (g *Gateway) Get(req GetRequest) {
	g.getChan <- req
}
