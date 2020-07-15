// 在这个例子中，我们将仅允许通过身份验证的用户在 /secret 中的一定有效期内
// 去查看我们的秘密消息。想要获得访问权限，首先必须去访问 /login 以获取有
// 效的会话 Cookie ，然后将通过验证的用户设置为登录状态。另外，他可以访问
// /logout 来撤销对我们秘密信息的访问

package main

import (
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(ctx iris.Context) {
	// 检查用户是否已通过身份验证
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	// 打印秘密消息
	ctx.WriteString("The cake is a lie!")
}

func login(ctx iris.Context) {
	log.Println("handle login")
	session := sess.Start(ctx)

	// 在此处进行身份验证
	// ...

	// 将用户设置为已验证
	session.Set("authenticated", true)
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)

	// 撤销用户身份验证
	session.Set("authenticated", false)
}

func main() {
	app := iris.New()

	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)

	app.Run(iris.Addr(":8080"))
}

// $ go run sessions.go

// $ curl -s http://localhost:8080/secret
// Forbidden

// $ curl -s -I http://localhost:8080/login
// Set-Cookie: mysessionid=MTQ4NzE5Mz...

// $ curl -s --cookie "mysessionid=MTQ4NzE5Mz..." http://localhost:8080/secret
// The cake is a lie!
