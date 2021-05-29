package main

import (
	"context"
	"fmt"
	pb "go-demos/grpc-demo/grpc-go/proto/echo"
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr = "localhost:50051"

type server struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *server) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{addr: addr})
	log.Printf("serving on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
