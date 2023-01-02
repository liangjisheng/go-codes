package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/client"
)

var (
	addr    = flag.String("addr", "localhost:8972", "server address")
	xclient client.XClient
)

func main() {
	flag.Parse()

	go http.ListenAndServe("127.0.0.1:9099", nil)
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient = client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
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
