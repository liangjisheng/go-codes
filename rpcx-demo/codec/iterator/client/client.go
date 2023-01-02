package main

import (
	"context"
	"flag"
	"log"

	"rpcx-demo/codec/iterator/codec"
	"rpcx-demo/service"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	// register customized codec
	share.Codecs[protocol.SerializeType(4)] = &codec.JsoniterCodec{}
	option := client.DefaultOption
	option.SerializeType = protocol.SerializeType(4)

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &service.Args{
		A: 10,
		B: 20,
	}

	reply := &service.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
