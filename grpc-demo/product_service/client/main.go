package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "go-demos/grpc-demo/product_service/product"
)

const (
	address = "localhost:5230"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	response, err := client.AddProduct(ctx, &pb.AddProductRequest{ProductName: "car"})
	if nil != err {
		log.Fatalf("add product failed, %v\n", err)
	}
	log.Printf("add product success, %s", response)

	productID := response.ProductID
	queryResp, err := client.QueryProductInfo(ctx, &pb.QueryProductRequest{ProductID: productID})
	if nil != err {
		log.Fatalf("query product info failed, %v\n", err)
	}
	log.Printf("Product info is %v\n", queryResp)

	defer cancel()
}
