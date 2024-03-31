package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// VisitController 处理根路由
type VisitController struct {
	// 当前请求的会话 通过我们添加到 `visitApp` 的依赖函数完成初始化
	Session *sessions.Session
	// 在 Controller 中绑定了 time.time，字段的顺序无关紧要
	StartTime time.Time
}

// Get 处理方法
// Method: GET
// Path: http://localhost:8080
func (c *VisitController) Get() string {
	// 每次访问，自增 "visits"
	// 如果 "visits"  不存在，将会创建一个
	visits := c.Session.Increment("visits", 1)
	// 使用当前更新过的 visits
	since := time.Now().Sub(c.StartTime).Seconds()
	return fmt.Sprintf("%d visit from my current session in %0.1f seconds of server's up-time",
		visits, since)
}

func newApp() *iris.Application {
	app := iris.New()
	sess := sessions.New(sessions.Config{Cookie: "mysession_cookit_name"})

	visitApp := mvc.New(app.Party("/"))
	// 绑定当前的 *session.Session 到 `VisitController.Session`，它是必要的
	// 并且添加 time.Now() 到 `VisitController.StartTime`
	visitApp.Register(
		// 如果依赖是接受 Context 并且返回单只的函数
		// 这个函数的结果将被控制器解析。
		// 当每个请求到来的时候，它将调用这个带有上下文的函数
		// 并且设置结果（这里是 *sessions.Session ）到控制器字段。
		//
		// 如果依赖项注册时没有字段或函数的输入参数作为使用者，
		// 那么这些依赖项将在服务器运行之前被忽略，
		// 因此您可以绑定许多依赖项并在不同的控制器中使用它们
		sess.Start,
		time.Now(),
	)
	visitApp.Handle(new(VisitController))
	return app
}

func main() {
	app := newApp()

	// 1. 打开浏览器（非私有模式）
	// 2. 打开网址： http://localhost:8080
	// 3. 刷新页面几次
	// 4. 关掉浏览器
	// 5. 重新打开浏览器，并且重复步骤2
	app.Run(iris.Addr(":8080"))
}
