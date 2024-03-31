package main

import (
	"fmt"
	"os"

	"github.com/opentracing/opentracing-go/log"

	"jaeger/opentracing-tutorial/lib"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := lib.InitJeager("hello1")
	defer closer.Close()

	span := tracer.StartSpan("say-hello1")

	helloTo := os.Args[1]
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	println(helloStr)
	span.LogKV("evnet", "println")

	span.Finish()
}
