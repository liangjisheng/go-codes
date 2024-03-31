package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// 当你想更变一些配置项的时候这样做
	// 前缀："With"，代码比较器将帮助你查看所有可用配置

	app.Run(iris.Addr(":8080"), iris.WithoutStartupLog, iris.WithCharset("UTF-8"))

	// or before run:
	// app.Configure(iris.WithoutStartupLog, iris.WithCharset("UTF-8"))
	// app.Run(iris.Addr(":8080"))
}
