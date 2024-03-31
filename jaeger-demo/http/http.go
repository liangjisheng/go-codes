package zk_tracing

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var span opentracing.Span
		tracer := opentracing.GlobalTracer()

		// if request header include span return child startSpan, else return parent startSpan
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			span = tracer.StartSpan(ctx.Request.URL.Path)
			span.SetTag(string(ext.Component), "HTTP")
			span.SetTag("Method", ctx.Request.Method)
		} else {
			span = opentracing.StartSpan(
				ctx.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "Router"},
				ext.SpanKindRPCServer,
			)
		}

		span.SetTag("Type", "request")

		ctx.Set("ParentSpanContext", span.Context())

		ctx.Next()
		span.Finish()
	}
}

func StartSpanHttp(ctx *gin.Context, callFuncName string) (opentracing.Span, context.Context) {
	parentSpanContext, _ := ctx.Get("ParentSpanContext")

	span := opentracing.StartSpan(
		callFuncName,
		opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
		opentracing.Tag{Key: string(ext.Component), Value: "Http"},
		ext.SpanKindRPCClient,
	)

	return span, opentracing.ContextWithSpan(context.Background(), span)
}
