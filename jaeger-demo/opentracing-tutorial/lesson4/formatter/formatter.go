package main

import (
	"fmt"
	"log"
	"net/http"

	"go-demos/jaeger-demo/opentracing-tutorial/lib"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := lib.InitJeager("formatter")
	defer closer.Close()

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		greeting := span.BaggageItem("greeting")
		if "" == greeting {
			greeting = "hello"
		}

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("%s, %s", greeting, helloTo)
		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", helloStr),
		)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
