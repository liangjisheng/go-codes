package main

import (
	"context"
	"fmt"

	proto "go-demos/micro-demos/getip/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

// Greeter ...
type Greeter struct{}

// Hello ...
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	md, _ := metadata.FromContext(ctx)

	// local ip of service
	fmt.Println("local ip is", md["Local"])

	// remote ip of caller
	fmt.Println("remote ip is", md["Remote"])

	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
