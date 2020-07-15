package main

import (
	"context"
	"fmt"
	proto "go-demos/micro-demos/getip/proto"

	"github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	fmt.Println("")
	// Create new greeter client
	client := proto.NewGreeterService("greeter", service.Client())

	rsp, err := client.Hello(context.TODO(), &proto.HelloRequest{Name: "ljs"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Greeting)
}
