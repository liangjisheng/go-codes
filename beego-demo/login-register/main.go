package main

import (
	_ "go-demos/beego-demo/login-register/models"
	_ "go-demos/beego-demo/login-register/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
