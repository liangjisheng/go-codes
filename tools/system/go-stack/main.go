package main

import (
	"github.com/go-stack/stack"
	"log"
)

func main() {
	c := stack.Caller(0)
	log.Print(c)         // "source.go:10"
	log.Printf("%+v", c) // "pkg/path/source.go:10"
	log.Printf("%n", c)  // "DoTheThing"

	s := stack.Trace().TrimRuntime()
	log.Print(s) // "[source.go:15 caller.go:42 main.go:14]"
}
