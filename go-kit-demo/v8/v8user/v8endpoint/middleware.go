package v8endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"v8/v8user/v8service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// LoggingMiddleware ...
func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v8service.ContextReqUUID)), zap.Any("调用 v8_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
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
