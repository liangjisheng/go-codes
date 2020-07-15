package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	app.Get("/", index)
	app.Run(iris.Addr(":8080"))
}

func notFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

func index(ctx iris.Context) {
	ctx.View("index.html")
}
