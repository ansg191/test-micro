package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	live "github.com/ansg191/test-micro/live/proto"
)

type Live struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Live) Call(ctx context.Context, req *live.Request, rsp *live.Response) error {
	log.Info("Received Live.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Live) Stream(ctx context.Context, req *live.StreamingRequest, stream live.Live_StreamStream) error {
	log.Infof("Received Live.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&live.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Live) PingPong(ctx context.Context, stream live.Live_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&live.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
