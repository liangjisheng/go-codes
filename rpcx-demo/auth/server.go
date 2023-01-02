package main

import (
	"context"
	"errors"
	"flag"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(service.Arith), "")
	s.AuthFunc = auth
	// s.Serve("reuseport", *addr)
	s.Serve("tcp", *addr)
}

func auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == "bearer tGzv3JOkF0XG5Qx2TlKWIA" {
		return nil
	}

	return errors.New("invalid token")
}
