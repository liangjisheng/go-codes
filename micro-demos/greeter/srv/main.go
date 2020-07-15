package main

import (
	"context"
	"log"
	"time"

	"github.com/micro/go-micro"

	pb "go-demos/micro-demos/greeter/srv/proto"
)

// Say ...
type Say struct{}

// Hello ...
func (s *Say) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	pb.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
