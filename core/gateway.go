package core

import grpc "./github.com/SoulTch/gomars_go"

// ActionRequestJob .
type ActionRequestJob struct {
	req grpc.ActionRequest
	res chan grpc.ActionResponse
}

// Gateway .
type Gateway struct {
	game       *Game
	actionChan chan ActionRequestJob
	getChan    chan grpc.GetRequest
	noticeChan chan grpc.NoticeResponse
}

// CreateGateway .
func CreateGateway(game *Game) *Gateway {
	gateway := new(Gateway)
	gateway.game = game
	gateway.actionChan = make(chan ActionRequestJob)
	gateway.getChan = make(chan grpc.GetRequest)
	gateway.noticeChan = make(chan grpc.NoticeResponse)

	return gateway
}

// Action .
func (g *Gateway) Action(req grpc.ActionRequest) grpc.ActionResponse {
	res := make(chan grpc.ActionResponse, 1)
	reqJob := ActionRequestJob{req, res}
	g.actionChan <- reqJob
	return <-res
}

// Notice .
func (g *Gateway) Notice(noti grpc.NoticeResponse) {
	g.noticeChan <- noti
}

// Get .
func (g *Gateway) Get(req grpc.GetRequest) {
	g.getChan <- req
}
