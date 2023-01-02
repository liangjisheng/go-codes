package main

import (
	"flag"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(service.Arith), "")
	s.Serve("tcp", *addr)
}
