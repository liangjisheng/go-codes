package main

import (
	"log"
	"net"

	v1 "grpcgateway/helloworld/v1"
	//v1 "grpcgateway/helloworld/v1buf"
	"grpcgateway/service"

	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":18080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	server := grpc.NewServer()
	// Attach the Greeter service to the server
	service := service.NewService()
	v1.RegisterGreeterServer(server, service)
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:18080")
	log.Fatal(server.Serve(lis))
}
