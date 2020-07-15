package main

import (
	common "demo/conf"
	"demo/utility/db"
	"demo/utility/helper"
	"demo/utility/log"
	"demo/proto/srv.user"
	
	"demo/app/srv.user/conf"
	"demo/app/srv.user/service"
	"demo/app/srv.user/handler"
	
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var appName = common.APP_SRV_USER

func main() {
	defer uninit()

	service.InitService(appName, micro.Action(initialize))
	_ = proto.RegisterLoginHandler(service.Server(), new(handler.LoginServer))
	helper.CheckErr("ServerRun", service.Run(), true)
}

func initialize(ctx *cli.Context) {
	// 初始化公共配置文件
	helper.CheckErr("InitCommonConfig", common.InitConfig(ctx.String("etcd_addr")), true)

	// 初始化app配置文件
	helper.CheckErr("InitAppConfig", conf.InitConfig(ctx.String("etcd_addr"), appName), true)

	// 初始化日志
	helper.CheckErr("InitZapLog", log.InitZapLogger(conf.GetLogPath()), true)

	// 启动mysql
	helper.CheckErr("InitMysql", db.InitMysql(common.GetMysqlConfig()), true)

	// 启动redis
	helper.CheckErr("InitRedis", db.InitRedis(common.GetRedisConfig()), true)
}

func uninit() {
	db.CloseMysql()
	db.CloseRedis()
}
