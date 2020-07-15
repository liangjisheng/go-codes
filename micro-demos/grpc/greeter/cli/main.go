package main

import (
	"context"
	"fmt"

	hello "go-demos/micro-demos/grpc/greeter/srv/proto"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro/metadata"
)

func main() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Hello(ctx, &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
