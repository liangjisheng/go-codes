package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	//　从　"./views"　目录下加载扩展名是".html" 　的所有模板，
	//　并使用标准的　`html/template`　 包进行解析
	app.RegisterView(iris.HTML("./views", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./views/hello.html
		ctx.View("hello.html")
	})

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	app.Get("/user/{id:long}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Run(iris.Addr(":8080"))
}
