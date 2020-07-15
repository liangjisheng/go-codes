package main

import (
	"go-demos/micro-demos/go-micro-demo/src/share/config"
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
	"go-demos/micro-demos/go-micro-demo/src/share/utils/log"
	"go-demos/micro-demos/go-micro-demo/src/user-srv/db"
	"go-demos/micro-demos/go-micro-demo/src/user-srv/handler"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

func main() {
	logger := log.Init("user")
	service := micro.NewService(
		micro.Name(config.Namespace+"user"),
		micro.Version("latest"),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("micro.Action test ...")
			// 先注册db
			db.Init(config.MysqlDSN)
			pb.RegisterUserServiceHandler(service.Server(), handler.NewUserHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("micro.AfterStop test ...")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("micro.AfterStart test ...")
			return nil
		}),
	)

	logger.Info("启动user-srv服务 ...")

	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败 ...")
	}
}
