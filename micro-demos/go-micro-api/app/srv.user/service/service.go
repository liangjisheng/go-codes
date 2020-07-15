package service

import (
	"context"
	"demo/utility/log"
	"fmt"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/valyala/fasthttp"
)

var microService micro.Service

// Run ...
func Run() error {
	return microService.Run()
}

// Server ...
func Server() server.Server {
	return microService.Server()
}

// NewContext ...
func NewContext(ctx *fasthttp.RequestCtx) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{
		"Request-Id": ctx.UserValue("Request-Id").(string),
	})
}

// NewClient ...
func NewClient() client.Client {
	return microService.Client()
}

func recoverHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		defer func() {
			if err := recover(); err != nil {
				var trace string
				for i := 1; ; i++ {
					if _, f, l, got := runtime.Caller(i); !got {
						break
					} else if strings.Contains(f, "/app/") {
						trace += fmt.Sprintf("%s:%d;", f, l)
					}
				}

				log.Error("recover exception.",
					zap.String("server", req.Endpoint()),
					zap.Any("reqBody", req.Body()),
					zap.Any("error", err),
					zap.String("trace", trace))
			}
		}()
		return fn(ctx, req, rsp)
	}
}

func accessHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		begin := time.Now()

		meta, _ := metadata.FromContext(ctx)

		defer func() {
			log.Info("access log.",
				zap.String("ip", meta["Ip-Addr"]),
				zap.String("method", "RPC"),
				zap.String("path", meta["Micro-Service"]+"."+meta["Micro-Method"]),
				zap.String("reqId", meta["Request-Id"]),
				zap.String("queries", ""),
				zap.Any("reqBody", req.Body()),
				zap.Any("respBody", rsp),
				zap.Duration("duration", time.Now().Sub(begin)))
		}()
		return fn(ctx, req, rsp)
	}
}
