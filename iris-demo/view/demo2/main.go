package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// - standard html  | iris.HTML(...)
	// - django         | iris.Django(...)
	// - pug(jade)      | iris.Pug(...)
	// - handlebars     | iris.Handlebars(...)
	// - amber          | iris.Amber(...)
	tmpl := iris.HTML("./templates", ".html")

	// 内置模板函数是：
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // 当前页面的部分相对路径
	// - {{ yield }}
	// - {{ current }}

	// 注册一个自定义模板函数
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	// 在视图中注册模板引擎，这样将会加载模板
	app.RegisterView(tmpl)

	app.Get("/", hi)
	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func hi(ctx iris.Context) {
	// 渲染模板文件　"./templates/hi.html"
	ctx.View("hi.html")
}
