package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", before, mainHandler, after)
	app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"
	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)
	ctx.Values().Set("info", shareInformation)
	ctx.Next() // 执行下一个处理器
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")
	// 获取 "before" 处理器中的设置的 "info" 值
	info := ctx.Values().GetString("info")

	// 响应客户端
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}
