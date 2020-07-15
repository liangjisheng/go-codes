package main

import (
	"context"
	"fmt"

	proto "go-demos/micro-demos/helloworld/proto"

	"github.com/micro/go-micro"
)

type greeter struct{}

func (g *greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	res.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
