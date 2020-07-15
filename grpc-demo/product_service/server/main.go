package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "go-demos/grpc-demo/product_service/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5230"
)

var dataBase = make(map[string]*Product, 10)

// Product ...
type Product struct {
	ProductName    string
	ProductID      string
	ManufacturerID string
	Weight         float64
	ProductionDate int64
	ImportDate     int64
}

type server struct{}

func (s *server) AddProduct(ctx context.Context, request *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	log.Printf("get request from client to add product, request is %s\n", request)
	productID := strconv.FormatInt(rand.Int63(), 10)
	product := new(Product)
	product.ProductName = request.ProductName
	product.ProductID = productID
	product.ManufacturerID = request.ManufacturerID
	product.Weight = request.Weight
	product.ProductionDate = request.ProductionDate
	product.ImportDate = time.Now().UnixNano()
	dataBase[productID] = product
	return &pb.AddProductResponse{ProductID: productID, Message: "Add product success"}, nil
}

func (s *server) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*pb.EmptyResponse, error) {
	log.Printf("get request from client to delete product, request is %s\n", request)
	productID := request.ProductID
	delete(dataBase, productID)
	return nil, nil
}

func (s *server) QueryProductInfo(ctx context.Context, request *pb.QueryProductRequest) (*pb.ProductInfoResponse, error) {
	log.Printf("get request from client for query product info,%v\n", request)
	productID := request.ProductID
	product := dataBase[productID]
	response := new(pb.ProductInfoResponse)
	response.ProductID = product.ProductID
	response.ProductName = product.ProductName
	response.ManufacturerID = product.ManufacturerID
	response.Weight = product.Weight
	response.ProductionDate = product.ProductionDate
	response.ImportDate = product.ImportDate
	return response, nil
}

func (s *server) QueryProductsInfo(ctx context.Context, request *pb.EmptyRequest) (*pb.ProductsInfoResponse, error) {
	log.Printf("get request from client fro query products info.\n")
	if len(dataBase) == 0 {
		return nil, errors.New("no product")
	}

	response := new(pb.ProductsInfoResponse)

	for productID, product := range dataBase {
		productInfoResponse := new(pb.ProductInfoResponse)
		productInfoResponse.ProductID = productID
		productInfoResponse.ProductName = product.ProductName
		productInfoResponse.ManufacturerID = product.ManufacturerID
		productInfoResponse.Weight = product.Weight
		productInfoResponse.ProductionDate = product.ProductionDate
		productInfoResponse.ImportDate = product.ImportDate
		response.Infos = append(response.Infos, productInfoResponse)
	}

	return response, nil
}

func main() {
	log.Printf("begin to start rpc server.\n")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
