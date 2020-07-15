package v9endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"v9/v9user/v9service"

	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// LoggingMiddleware ...
func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v9service.ContextReqUUID)), zap.Any("调用 v9_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// NewGolangRateAllowMiddleware ...
func NewGolangRateAllowMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return "", errors.New("limit req allow")
			}
			return next(ctx, request)
		}
	}
}

// NewTracerEndpointMiddleware ...
func NewTracerEndpointMiddleware(tracer opentracing.Tracer) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "endpoint", opentracing.Tag{
				Key:   string(ext.Component),
				Value: "NewTracerEndpointMiddleware",
			})
			defer span.Finish()
			return next(ctxContext, request)
		}
	}
}
