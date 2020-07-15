package service

import (
	"context"
	"demo/utility/helper"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/valyala/fasthttp"
)

var microService micro.Service

// InitService ...
func InitService(appName string, opts ...micro.Option) {
	microService = micro.NewService(
		micro.Name(appName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Flags(cli.StringFlag{
			Name:   "etcd_addr",
			EnvVar: "ETCD_ADDR",
			Usage:  "This is etcd config address.",
		}),
	)

	microService.Init(opts...)
}

// NewContext ...
func NewContext(ctx *fasthttp.RequestCtx) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{
		"Request-Id": ctx.UserValue("Request-Id").(string),
		"Ip-Addr":    string(helper.RealIPAddr(ctx)),
	})
}

// NewClient ...
func NewClient() client.Client {
	return microService.Client()
}
