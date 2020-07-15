package v4endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"v4/utils"
	"v4/v4service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// LoggingMiddleware ...
func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v4service.ContextReqUUID)), zap.Any("调用 v3_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// AuthMiddleware ...
func AuthMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(utils.JwtContextKey))
			if token == "" {
				err = errors.New("请登录")
				logger.Debug(fmt.Sprint(ctx.Value(v4service.ContextReqUUID)), zap.Any("[AuthMiddleware]", "token == empty"), zap.Error(err))
				return "", err
			}
			jwtInfo, err := utils.ParseToken(token)
			if err != nil {
				logger.Debug(fmt.Sprint(ctx.Value(v4service.ContextReqUUID)), zap.Any("[AuthMiddleware]", "ParseToken"), zap.Error(err))
				return "", err
			}
			if v, ok := jwtInfo["Name"]; ok {
				ctx = context.WithValue(ctx, interface{}("name"), v)
			}
			return next(ctx, request)
		}
	}
}

// NewGolangRateWaitMiddleware ...
func NewGolangRateWaitMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err = limit.Wait(ctx); err != nil {
				return "", errors.New("limit req  Wait")
			}
			return next(ctx, request)
		}
	}
}

// NewGolangRateAllowMiddleware ...
func NewGolangRateAllowMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return "", errors.New("limit req  Allow")
			}
			return next(ctx, request)
		}
	}
}

// NewUberRateMiddleware ...
func NewUberRateMiddleware(limit ratelimit.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			limit.Take()
			return next(ctx, request)
		}
	}
}
