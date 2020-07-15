package main

import (
	"go-demos/beego-demo/user-register-login/controllers"
	"go-demos/beego-demo/user-register-login/models"
	_ "go-demos/beego-demo/user-register-login/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()        // 数据库初始化
	models.InitSession() // 初始化session配置
}

func main() {
	// 过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, controllers.BeforeExecFilter, false, false)
	beego.InsertFilter("/user/*", beego.BeforeRouter, controllers.BeforeRouterFilter, false, false)
	beego.InsertFilter("/user/*", beego.BeforeStatic, controllers.BeforeStaticFilter, false, false)
	beego.InsertFilter("/user/*", beego.AfterExec, controllers.AfterExecFilter, false, false)
	beego.InsertFilter("/user/*", beego.FinishRouter, controllers.FinishRouterFilter, false, false)

	// run方法前注册错误handler
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}
