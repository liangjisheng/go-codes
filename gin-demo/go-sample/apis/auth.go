package apis

import (
	"fmt"
	"go-demos/gin-demo/go-sample/models"
	"go-demos/gin-demo/go-sample/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
// url中使用&传递多参数，在linux系统中 &会使进程系统后台运行
// 1.可以对&转义 \&
// curl -s -X GET http://127.0.0.1:8080/login?name=alice\&password=123456
// 用双引号把整个url引起来
// curl -X GET "http://127.0.0.1:8080/login?name=alice&password=123456"
func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	// fmt.Println("name: ", name)
	// fmt.Println("password: ", password)

	if hasSession := pkg.HasSession(c); hasSession == true {
		c.String(200, "用户已登陆")
		return
	}

	user := models.UserDetailByName(name)
	// fmt.Printf("user: %+v", user)
	if err := pkg.Compare(user.Password, password); err != nil {
		c.String(401, "密码错误")
		return
	}

	pkg.SaveAuthSession(c, user.ID)
	// session := sessions.Default(c)
	// if sessionValue := session.Get("userId"); sessionValue == nil {
	// 	fmt.Println("session value is nil")
	// } else {
	// 	fmt.Println("session value:", sessionValue)
	// }
	c.String(200, "登录成功")
}

// Logout ...
// curl http://127.0.0.1:8080/logout
func Logout(c *gin.Context) {
	uid := pkg.GetSessionUserID(c)
	fmt.Println("user id:", uid)

	if hasSession := pkg.HasSession(c); hasSession == false {
		c.String(401, "用户未登录")
		return
	}
	pkg.ClearAuthSession(c)
	c.String(200, "退出成功")
}

// Register ...
// curl http://127.0.0.1:8080/register -X POST -d 'name=alice&email=example@qq.com&password=123456&password_confirmation=123456'
func Register(c *gin.Context) {
	var user models.User
	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")

	if hasSession := pkg.HasSession(c); hasSession == true {
		c.String(200, "用户已登陆")
		return
	}

	if existUser := models.UserDetailByName(user.Name); existUser.ID != 0 {
		c.String(200, "用户名已存在")
		return
	}

	if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
		c.String(200, "密码不一致")
		return
	}

	if pwd, err := pkg.Encrypt(c.Request.FormValue("password")); err == nil {
		user.Password = pwd
	}

	models.AddUser(&user)
	pkg.SaveAuthSession(c, user.ID)
	c.String(200, "注册成功")
}

// Me ...
// curl http://127.0.0.1:8080/auth/me
func Me(c *gin.Context) {
	currentUser := c.MustGet("userId").(uint)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": currentUser,
	})
}
