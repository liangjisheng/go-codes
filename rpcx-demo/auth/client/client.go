package main

import (
	"context"
	"flag"
	"log"
	"rpcx-demo/service"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	option := client.DefaultOption
	//option.ReadTimeout = 10 * time.Second

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	// xclient.Auth("bearer tGzv3JOkF0XG5Qx2TlKWIA")
	xclient.Auth("bearer abcdefg1234567")

	args := &service.Args{
		A: 10,
		B: 20,
	}

	reply := &service.Reply{}
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
