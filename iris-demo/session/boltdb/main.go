// 有时你需要一个后端的存储，即文件存储或者 Redis 内存数据库存储，这可以让
// 你的会话数据在服务器重启后保持不变
// 通过调用 .UseDatabase(database) 来注册一个数据库是非常容易的

package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
)

func main() {
	db, _ := boltdb.New("./sessions.db", 0666)
	// 使用不同的协程来同步数据库
	// db.Async(true)

	// 按下 control+C/cmd+C 来关闭并解锁数据库
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	sess := sessions.New(sessions.Config{
		Cookie:  "sessionscookieid",
		Expires: 45 * time.Minute, // <=0 意味永久的存活
	})

	//
	// 非常重要：
	//
	sess.UseDatabase(db)

	// 剩下的其它代码保持不变。
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})

	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx)
		// 设置一个 session 值
		s.Set("name", "iris")

		// 在这测试已设置的 session 值
		ctx.Writef("All ok session setted to: %s", s.GetString("name"))
	})

	app.Get("/set/{key}/{value}", func(ctx iris.Context) {
		key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
		s := sess.Start(ctx)
		// 设置一个 session 值
		s.Set(key, value)

		// 在这测试已设置的 session 值
		ctx.Writef("All ok session setted to: %s", s.GetString(key))
	})

	app.Get("/get", func(ctx iris.Context) {
		// 获取一个特定的键，如字符串，如果没有获取到，则返回一个空字符串
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/get/{key}", func(ctx iris.Context) {
		// 获取一个特定的键，如字符串，如果没有获取到，则返回一个空字符串
		name := sess.Start(ctx).GetString(ctx.Params().Get("key"))

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx iris.Context) {
		// 删除一个特定的键
		sess.Start(ctx).Delete("name")
	})

	app.Get("/clear", func(ctx iris.Context) {
		// 删除所有键值对
		sess.Start(ctx).Clear()
	})

	app.Get("/destroy", func(ctx iris.Context) {
		// destroy方法,删除整个会话数据和 Cookie
		sess.Destroy(ctx)
	})

	app.Get("/update", func(ctx iris.Context) {
		// 更新过期的日期以及新的日期
		sess.ShiftExpiration(ctx)
	})

	app.Run(iris.Addr(":8080"))
}
