package v3endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"
	"v3/utils"

	"v3/v3service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// LoggingMiddleware ...
func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v3service.ContextReqUUID)), zap.Any("调用 v3_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
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
				logger.Debug(fmt.Sprint(ctx.Value(v3service.ContextReqUUID)), zap.Any("[AuthMiddleware]", "token == empty"), zap.Error(err))
				return "", err
			}
			jwtInfo, err := utils.ParseToken(token)
			if err != nil {
				logger.Debug(fmt.Sprint(ctx.Value(v3service.ContextReqUUID)), zap.Any("[AuthMiddleware]", "ParseToken"), zap.Error(err))
				return "", err
			}
			if v, ok := jwtInfo["Name"]; ok {
				ctx = context.WithValue(ctx, interface{}("name"), v)
			}
			return next(ctx, request)
		}
	}
}
