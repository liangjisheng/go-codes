package main

import (
	"context"
	"fmt"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"jaeger/opentracing-tutorial/lib"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := lib.InitJeager("hello2")
	defer closer.Close()

	rootSpan := tracer.StartSpan("say-hello2")

	helloTo := os.Args[1]
	rootSpan.SetTag("hello-to", helloTo)
	defer rootSpan.Finish()

	// case1
	// helloStr := formatString(rootSpan, helloTo)
	// printHello(rootSpan, helloStr)

	// case2
	// helloStr := formatString1(rootSpan, helloTo)
	// printHello1(rootSpan, helloStr)

	// case3
	// helloStr := formatString2(rootSpan, helloTo)
	// printHello2(rootSpan, helloStr)

	// case4
	opentracing.SetGlobalTracer(tracer)
	ctx := opentracing.ContextWithSpan(context.Background(), rootSpan)

	helloStr := formatString3(ctx, helloTo)
	printHello3(ctx, helloStr)
}

func formatString(span opentracing.Span, helloTo string) string {
	helloStr := fmt.Sprintf("hello, %s", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	return helloStr
}

func printHello(span opentracing.Span, helloStr string) {
	println(helloStr)
	span.LogKV("event", "println")
}

func formatString1(rootSpan opentracing.Span, helloTo string) string {
	span := rootSpan.Tracer().StartSpan("formatString1")
	defer span.Finish()

	helloStr := fmt.Sprintf("hello, %s", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	return helloStr
}

func printHello1(rootSpan opentracing.Span, helloStr string) {
	span := rootSpan.Tracer().StartSpan("printHello1")
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}

func formatString2(rootSpan opentracing.Span, helloTo string) string {
	span := rootSpan.Tracer().StartSpan(
		"formatString1",
		opentracing.ChildOf(rootSpan.Context()),
	)
	defer span.Finish()

	helloStr := fmt.Sprintf("hello, %s", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	return helloStr
}

func printHello2(rootSpan opentracing.Span, helloStr string) {
	span := rootSpan.Tracer().StartSpan(
		"printHello1",
		opentracing.ChildOf(rootSpan.Context()),
	)
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}

func formatString3(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString1")
	defer span.Finish()

	helloStr := fmt.Sprintf("hello, %s", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	return helloStr
}

func printHello3(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello1")
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}
