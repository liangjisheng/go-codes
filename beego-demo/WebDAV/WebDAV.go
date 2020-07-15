package main

import (
	"go-demos/beego-demo/WebDAV/controllers"

	"github.com/astaxie/beego"
)

func main() {
	// Register routers.
	beego.Router("/*", &controllers.WebDAVController{}, "*:Main")
	beego.Run()
}
