package main

import (
	"context"

	"github.com/micro/go-micro"
)

// Greeter ...
type Greeter struct{}

// Hello ...
func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// create new service
	srv := micro.NewService(
		micro.Name("greeter"),
	)

	// initialise command line
	srv.Init()

	// set the handler
	micro.RegisterHandler(srv.Server(), new(Greeter))

	// run service
	srv.Run()
}
