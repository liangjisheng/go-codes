// Package service defines datastructure and services.
package service

import (
	"context"
	"fmt"
)

// Args ...
type Args struct {
	A int
	B int
}

// Reply ..
type Reply struct {
	C int
}

// Arith ...
type Arith int

// Mul ...
func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

// Add ...
func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

// Say ...
func (t *Arith) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args
	return nil
}
