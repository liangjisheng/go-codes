package main

import (
	"flag"

	"rpcx-demo/codec/iterator/codec"
	"rpcx-demo/service"

	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	share.Codecs[protocol.SerializeType(4)] = &codec.JsoniterCodec{}
	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(service.Arith), "")
	s.Serve("tcp", *addr)
}
