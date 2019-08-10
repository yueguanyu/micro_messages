package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	messages "github.com/yueguanyu/messages/proto/messages"
)

type Messages struct{}

func (e *Messages) Handle(ctx context.Context, msg *messages.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *messages.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
