package main

import (
	"context"
	"flag"
	"log"
	"rpcx-demo/service"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	option := client.DefaultOption
	//option.ReadTimeout = 10 * time.Second

	xclient := client.NewXClient("a.b.c.D", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &service.Args{
		A: 10,
		B: 20,
	}

	reply := &service.Reply{}
	err := xclient.Call(context.Background(), "Times", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
