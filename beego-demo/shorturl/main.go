package main

import (
	"go-demos/beego-demo/shorturl/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})
	beego.Run()
}
