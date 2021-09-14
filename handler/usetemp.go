package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	usetemp "usetemp/proto"
)

type Usetemp struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Usetemp) Call(ctx context.Context, req *usetemp.Request, rsp *usetemp.Response) error {
	log.Info("Received Usetemp.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Usetemp) Stream(ctx context.Context, req *usetemp.StreamingRequest, stream usetemp.Usetemp_StreamStream) error {
	log.Infof("Received Usetemp.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&usetemp.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Usetemp) PingPong(ctx context.Context, stream usetemp.Usetemp_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&usetemp.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}