package main

// https://hacpai.com/article/1564238427681

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"jaeger/demo/tracer"
	"log"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

// curl http://127.0.0.1:8002/getIP

func main() {
	var (
		err error
		io  io.Closer
	)

	tracer.Tracer, io, err = tracer.NewTracer("serviceName", "117.51.148.112:6831")
	if err != nil {
		log.Fatalf("tracer.NewTracer error(%v)", err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer.Tracer)

	http.HandleFunc("/getIP", getIP)
	log.Printf("Starting server on port %d", 8002)
	err = http.ListenAndServe("127.0.0.1:8002",
		nethttp.Middleware(tracer.Tracer, http.DefaultServeMux))
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
}

func getIP(w http.ResponseWriter, r *http.Request) {
	log.Print("Received getIP request")

	// client
	client := &http.Client{Transport: &nethttp.Transport{}}
	span := tracer.Tracer.StartSpan("getIP")
	span.SetTag(string(ext.Component), "getIP")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	req, err := http.NewRequest("GET", "http://icanhazip.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)
	// warp the request in nethttp.TraceRequest
	req, ht := nethttp.TraceRequest(tracer.Tracer, req)
	defer ht.Finish()

	res, err := client.Do(req)
	if err != nil {
		onError(span, err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		onError(span, err)
		return
	}
	log.Printf("Received result: %s\n", body)
	io.WriteString(w, fmt.Sprintf("ip %s", body))
}

func onError(span opentracing.Span, err error) {
	span.SetTag(string(ext.Error), true)
	span.LogKV(otlog.Error(err))
	log.Fatalf("client(%v)", err)
}
