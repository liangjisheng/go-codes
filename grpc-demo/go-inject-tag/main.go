package main

import (
	"encoding/json"
	"fmt"
	"go-demos/grpc-demo/go-inject-tag/inject"
)

func main() {
	msg := inject.MyMessage{
		Code: 0,
	}

	data, _ := json.Marshal(msg)
	fmt.Println(string(data))
}
