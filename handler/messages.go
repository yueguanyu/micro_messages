package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	messages "github.com/yueguanyu/messages/proto/messages"
)

type Messages struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Messages) Call(ctx context.Context, req *messages.Request, rsp *messages.Response) error {
	log.Log("Received Messages.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Messages) Stream(ctx context.Context, req *messages.StreamingRequest, stream messages.Messages_StreamStream) error {
	log.Logf("Received Messages.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&messages.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Messages) PingPong(ctx context.Context, stream messages.Messages_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&messages.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
