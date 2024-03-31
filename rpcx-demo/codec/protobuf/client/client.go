package main

import (
	"context"
	"flag"
	"log"

	"rpcx-demo/codec/protobuf/pb"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	// register customized codec
	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &pb.ProtoArgs{
		A: 10,
		B: 20,
	}

	reply := &pb.ProtoReply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
