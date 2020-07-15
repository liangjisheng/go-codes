package main

import (
	_ "go-demos/beego-demo/WebIM/routers"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	APP_VER = "0.1.1.0.227"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)

	// Register tmeplate functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
