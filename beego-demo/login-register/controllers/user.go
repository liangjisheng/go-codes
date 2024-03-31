package controllers

import (
	"fmt"
	"go-demos/beego-demo/login-register/models/class"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// UserController ...
type UserController struct {
	beego.Controller
}

// PageLogin ...
func (c *UserController) PageLogin() {
	c.TplName = "login.html" // 将login.html页面输出
}

// Register ...
func (c *UserController) Register() {
	id, _ := c.GetInt("userid")
	password := c.GetString("password")
	fmt.Println("This is id and password")
	fmt.Println(id, password)

	valid := validation.Validation{}
	valid.Required(id, "id")
	valid.Required(password, "password")

	switch {
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "bad.html"
		return
	}

	u := &class.User{
		ID:       id,
		Password: password,
	}

	err := u.Create()
	if err != nil {
		fmt.Println(err)
		c.TplName = "bad.html"
		return
	}
	c.TplName = "welcome.html"
}

// Reallogin ...
func (c *UserController) Reallogin() {
	id, _ := c.GetInt("userid")
	password := c.GetString("password")
	u := &class.User{
		ID:       id,
		Password: password,
	}

	err := u.ReadDB()
	if err != nil {
		fmt.Println(err)
		c.TplName = "bad.html"
		return
	}
	c.TplName = "welcome.html"
}
