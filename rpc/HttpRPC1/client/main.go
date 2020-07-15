package main

import (
	"log"
	"net/rpc"
	"os"
)

// Result ...
type Result struct {
	Num, Ans int
}

var client *rpc.Client

func init() {
	var err error
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(`rpc.DialHTTP("tcp", "127.0.0.1:8080") err:`, err)
		os.Exit(-1)
	}
}

func syncCall() {
	// client.Call 是同步调用的方式，会阻塞当前的程序，直到结果返回
	// 如果有异步调用的需求，可以考虑使用client.Go
	var res Result
	if err := client.Call("Cal.Square", 12, &res); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}
	log.Printf("%d^2 = %d", res.Num, res.Ans)
}

func asyncCall() {
	var res Result
	aCall := client.Go("Cal.Square", 12, &res, nil)
	log.Printf("%d^2 = %d", res.Num, res.Ans)

	<-aCall.Done
	log.Printf("%d^2 = %d", res.Num, res.Ans)
}

func main() {
	// syncCall()
	asyncCall()
}
