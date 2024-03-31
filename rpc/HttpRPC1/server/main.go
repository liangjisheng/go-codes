package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// Result ...
type Result struct {
	Num, Ans int
}

// Cal ...
type Cal int

// Square ...
func (cal *Cal) Square(num int, res *Result) error {
	res.Num = num
	res.Ans = num * num
	return nil
}

func main() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 8080)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal("Error serving ", err)
	}
}
