package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// 当你有两份不同的配置时这样使用，一份用于开发，一份用于生产
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
}
