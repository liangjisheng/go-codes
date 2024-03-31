package main

// https://blog.csdn.net/liyunlong41/article/details/87932953

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func initJaeger(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		// 此时我们要在reporter中配置jaeger Agent的ip与端口，以便将tracer的信息发布到agent中
		// 配置LocalAgentHostPort参数为127.0.0.1:6381，6381接口是接受压缩格式的thrift协议数据
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "117.51.148.112:6831",
		},
	}

	tracer, closer, err := cfg.New(serviceName, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func foo(ctx context.Context, req string) (reply string) {
	// 创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo")
	defer func() {
		// 接口调用完,在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()

	println(req)
	// 模拟处理耗时
	time.Sleep(2 * time.Second)
	reply = "fooReply"
	return
}

func foo1(ctx context.Context, req string) (reply string) {
	// 创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo1")
	defer func() {
		// 接口调用完,在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()

	println(req)
	// 模拟处理耗时
	time.Sleep(2 * time.Second)
	reply = "foo1Reply"
	return
}

func main() {
	tracer, closer := initJaeger("demo1")
	defer closer.Close()

	// StartSpanFromContext会用到opentracing.SetGlobalTracer()
	// 来启动新的span，所以在main函数中需要调用
	opentracing.SetGlobalTracer(tracer)
	// 创建一个root span，调用两个函数，分别表示调用两个分布式服务
	span := tracer.StartSpan("span_root")
	// 用ContextWithSpan来创建一个新的ctx，将span的信息与context关联
	// 传到foo中时，需要创建一个子span，父span是ctx中的span
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	r1 := foo(ctx, "hello foo")
	r2 := foo1(ctx, "hello foo1")
	fmt.Println(r1, r2)
	span.Finish()
}
