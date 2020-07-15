package main

import (
	"context"
	"fmt"
	pb "go-demos/micro-demos/greeter/srv/proto"

	"github.com/micro/go-micro"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := pb.NewSayService("go.micro.srv.greeter", service.Client())

	// Make request
	rsp, err := cl.Hello(context.Background(), &pb.Request{
		Name: "ljs",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
