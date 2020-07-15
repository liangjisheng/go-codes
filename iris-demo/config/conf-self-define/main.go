package main

import (
	"go-demos/iris-demo/config/conf-self-define/counter"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Configure(counter.Configurator)

	app.Run(iris.Addr(":8080"))
}
