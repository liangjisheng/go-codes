package main

import (
	"context"
	"fmt"
	"time"

	proto "go-demos/micro-demos/helloworld/proto"

	"github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Creater new greeter
	ticker := time.NewTicker(time.Second)
	for {
		// Call the greeter
		rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "ljs"})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(rsp.Greeting)
		<-ticker.C
	}
}
