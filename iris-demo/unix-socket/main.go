package main

import (
	"github.com/kataras/iris"
	"github.com/valyala/tcplisten"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}

	l, err := listenerCfg.NewListener("tcp", ":8080")
	if err != nil {
		app.Logger().Fatal(err)
	}

	app.Run(iris.Listener(l))
}
