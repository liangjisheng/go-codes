package main

import (
	"context"
	"flag"
	"log"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	args := service.Args{
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
