package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"demo.book.com/conf"
	"demo.book.com/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
)

// curl 'http://127.0.0.1:8080/'
// curl 'http://127.0.0.1:8080/demo/q/p'
// curl 'http://127.0.0.1:8080/demo/x/m'
// curl 'http://127.0.0.1:8080/demo/record1'
// curl 'http://127.0.0.1:8080/demo/conf'
// curl 'http://127.0.0.1:8080/demo/err'
// curl 'http://127.0.0.1:8080/demo/orm'

// 浏览器访问: http://127.0.0.1:8080/book
// curl 'http://127.0.0.1:8080/book'

func main() {
	// 创建iris
	app := iris.New()
	// 设置debug模式
	app.Logger().SetLevel("debug")
	// 同时写文件日志与控制台日志
	app.Logger().SetOutput(io.MultiWriter(func() *os.File {
		// 创建日志目录
		dirName := "log/" + time.Now().Format("2006/01")
		os.MkdirAll(dirName, os.ModePerm)

		today := time.Now().Format(conf.SysTimeformShort)
		filename := dirName + "/" + today + ".log"
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		return f
	}(), os.Stdout))

	// 设置记录日志的格式中间件
	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},
		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	// 全局错误捕获中间件
	customRecover := func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}
					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				errMsg := fmt.Sprintf("错误信息: %s", err)
				// when stack finishes
				logMessage := fmt.Sprintf("从错误中回复：('%s')\n", ctx.HandlerName())
				logMessage += errMsg + "\n"
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				// 打印错误日志
				ctx.Application().Logger().Warn(logMessage)
				// 返回错误信息
				ctx.JSON(errMsg)
				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()
		ctx.Next()
	}
	app.Use(customRecover)

	// 注册模板
	app.RegisterView(iris.HTML("./web/views", ".html"))
	// 注册静态文件
	app.StaticWeb("/content", "./web/content")
	// app.HandleDir("/content", "./web/content")

	// 注册控制路由
	mvc.New(app.Party("/book")).Handle(new(controllers.BookController))
	mvc.New(app.Party("/demo")).Handle(new(controllers.DemoController))

	app.Run(
		iris.Addr(conf.SysConfMap["domain"]+":"+conf.SysConfMap["port"]),
		// 启动时禁止检测框架版本差异
		// iris.WithoutVersionChecker,
		// 忽略服务器错误
		iris.WithoutServerError(iris.ErrServerClosed),
		// 让程序自身尽可能的优化
		iris.WithOptimizations,
		iris.WithCharset("UTF-8"), // 国际化
	)
}
