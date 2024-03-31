package main

import (
	"context"
	"flag"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
)

// Arith ...
type Arith int

// Mul ...
func (t *Arith) Mul(ctx context.Context, args *service.Args, reply *service.Reply) error {
	reply.C = args.A * args.B * 100
	return nil
}

func main() {
	flag.Parse()

	go createServer1(*addr1, "")
	go createServer2(*addr2, "")

	select {}
}

func createServer1(addr, meta string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(service.Arith), meta)
	s.Serve("tcp", addr)
}

func createServer2(addr, meta string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), meta)
	s.Serve("tcp", addr)
}
