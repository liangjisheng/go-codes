package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"

	// 当然也可以不使用模版，直接用 c.Ctx.WriteString 输出字符串
	c.Ctx.WriteString("hello")
}
