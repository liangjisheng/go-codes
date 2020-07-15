package main

import (
	"context"
	proto "go-demos/micro-demos/greeter/api/rpc/hello"
	hello "go-demos/micro-demos/greeter/srv/proto"
	"log"

	"github.com/micro/go-micro"
)

// Greeter ...
type Greeter struct {
	Client hello.SayService
}

// Hello ...
func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Greeter.Hello API request")

	// make the request
	response, err := g.Client.Hello(ctx, &hello.Request{
		Name: req.Name,
	})
	if err != nil {
		return err
	}

	// set api response
	rsp.Msg = response.Msg
	return nil
}

func main() {
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	proto.RegisterGreeterHandler(service.Server(), &Greeter{
		Client: hello.NewSayService("go.micro.srv.greeter", service.Client()),
	})

	// for handler use

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
