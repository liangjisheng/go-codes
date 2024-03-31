package main

import (
	"bytes"
	"encoding/gob"
	"flag"

	"rpcx-demo/service"

	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	share.Codecs[protocol.SerializeType(4)] = &GobCodec{}
	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(service.Arith), "")
	s.Serve("tcp", *addr)
}

// GobCodec ...
type GobCodec struct {
}

// Decode ...
func (c *GobCodec) Decode(data []byte, i interface{}) error {
	enc := gob.NewDecoder(bytes.NewBuffer(data))
	err := enc.Decode(i)
	return err
}

// Encode ...
func (c *GobCodec) Encode(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)
	return buf.Bytes(), err
}
