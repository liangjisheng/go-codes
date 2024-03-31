package main

import (
	"context"
	"flag"
	"fmt"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "server address")
)

// Arith ...
type Arith struct{}

// Mul the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args service.Args, reply *service.Reply) error {
	reply.C = args.A * args.B
	fmt.Println("C=", reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
