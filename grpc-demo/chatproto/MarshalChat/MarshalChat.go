package main

import (
	"encoding/hex"
	"fmt"
	"go-demos/grpc-demo/chatproto/chat"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	var a = &chat.Request{Input: "alice"}
	protoA, err := proto.Marshal(a)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(hex.EncodeToString(protoA))

	var a1 chat.Request
	err = proto.Unmarshal(protoA, &a1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(a1.Input)
}
