package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/yueguanyu/messages/handler"
	"github.com/yueguanyu/messages/subscriber"

	messages "github.com/yueguanyu/messages/proto/messages"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.messages"),
		micro.Version("0.0.1"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	messages.RegisterMessagesHandler(service.Server(), new(handler.Messages))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.messages", service.Server(), new(subscriber.Messages))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.messages", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
