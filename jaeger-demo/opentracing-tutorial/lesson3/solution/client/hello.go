package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"jaeger/opentracing-tutorial/lib"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// go run hello.go alice

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := lib.InitJeager("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	helloTo := os.Args[1]

	rootSpan := tracer.StartSpan("say-hello")
	rootSpan.SetTag("hello-to", helloTo)
	defer rootSpan.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), rootSpan)

	helloStr := formatString(ctx, helloTo)
	printHello(ctx, helloStr)
}

func formatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloTo", helloTo)
	url := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := lib.Do(req)
	if err != nil {
		panic(err.Error())
	}

	helloStr := string(resp)

	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	return helloStr
}

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloStr", helloStr)
	url := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	if _, err := lib.Do(req); err != nil {
		panic(err.Error())
	}
}
