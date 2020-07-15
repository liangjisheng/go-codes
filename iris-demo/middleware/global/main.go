package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// 注册 "before"  处理器作为当前域名所有路由中第一个处理函数
	// 或者使用  `UseGlobal`  去注册一个中间件，用于在所有子域名中使用
	app.Use(before)
	// 注册  "after" 在所有路由的处理程序之后调用
	app.Done(after)

	// 注册路由
	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)

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

func indexHandler(ctx iris.Context) {
	ctx.HTML("<h1>Index</h1>") // 响应客户端
	ctx.Next()                 // 执行通过 `Done` 注册的 "after" 处理器
}

func contactHandler(ctx iris.Context) {
	ctx.HTML("<h1>Contact</h1>") // 响应客户端
	ctx.Next()                   // 执行通过 `Done` 注册的 "after" 处理器。
}
