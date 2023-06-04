package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "grpcgateway/helloworld/v1"
	//v1 "grpcgateway/helloworld/v1buf"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:18080"
	defaultName = "hello"
)

func sayHello(client v1.GreeterClient) {
	// Contact the serve and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &v1.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s\n", r.Message)
}

func hello1(client v1.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	in := &emptypb.Empty{}
	body, err := client.Hello1(ctx, in)
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Printf("body: %s\n", string(body.Data))
}

func download(client v1.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	in := &emptypb.Empty{}
	stream, err := client.Download(ctx, in)
	if err != nil {
		log.Println("err", err)
		return
	}

	for {
		body, err := stream.Recv()
		if err != nil && err != io.EOF {
			log.Println("err", err)
			return
		}

		if err == io.EOF {
			break
		}
		log.Println("body", string(body.Data))
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := v1.NewGreeterClient(conn)

	//sayHello(client)
	//hello1(client)
	download(client)
}
