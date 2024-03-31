package main

import (
	"flag"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	a := serverplugin.NewAliasPlugin()
	a.Alias("a.b.c.D", "Times", "Arith", "Mul")
	s := server.NewServer()
	s.Plugins.Add(a)
	s.RegisterName("Arith", new(service.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
