package main

import (
	"context"
	"fmt"
	v1 "grpcgateway/helloworld/v1"
	"grpcgateway/service"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	//v1 "grpcgateway/helloworld/v1buf"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:18080")
	go func() {
		log.Fatalln(server.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:18080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()

	// Attachment upload from http/s handled manually
	mux.HandlePath("POST", "/v1/files", handleBinaryFileUpload)

	// Register Greeter
	err = v1.RegisterGreeterHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":18090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:18090")
	log.Fatalln(gwServer.ListenAndServe())
}

func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	f, header, err := r.FormFile("attachment")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()
	log.Println("filename:", header.Filename)

	//
	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.
	//

	dst, err := os.OpenFile(header.Filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, f)
}
