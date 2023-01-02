package main

import (
	"context"
	"flag"
	"log"
	"rpcx-demo/service"
	"time"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &service.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &service.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		err = xclient.Call(context.Background(), "Add", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d + %d = %d", args.A, args.B, reply.C)

		time.Sleep(1e9)
	}
}
