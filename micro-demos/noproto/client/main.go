package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

func main() {
	srv := micro.NewService()
	srv.Init()
	c := srv.Client()

	req := c.NewRequest("greeter", "Greeter.Hello", "ljs", client.WithContentType("application/json"))
	var rsp string

	if err := c.Call(context.TODO(), req, &rsp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
