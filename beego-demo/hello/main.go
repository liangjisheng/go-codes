package main

import (
	_ "go-demos/beego-demo/hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 用户可以设置多个静态文件处理目录
	// 这样用户访问 URL http://localhost:8080/down1/123.txt 则会请求 download1 目录下的 123.txt 文件
	beego.SetStaticPath("/down1", "download1")
	beego.SetStaticPath("/down2", "download2")

	// beego 默认注册了 static 目录为静态处理的目录
	// 在/main.go文件中beego.Run()之前加入StaticDir["/static"] = "static"
	beego.Run()
}
