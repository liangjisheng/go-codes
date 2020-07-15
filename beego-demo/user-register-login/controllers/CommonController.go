package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// CommonController ...
type CommonController struct {
	beego.Controller
}

//URLMapping 配置注解路由
func (c *CommonController) URLMapping() {
	c.Mapping("Success", c.Success)
	c.Mapping("False", c.False)
}

// Success ...
// @router /success [get]
func (c *CommonController) Success() {
	logs.Info(">>>> forward to success page start <<<<")
	c.TplName = "success.html"
}

// False ...
// @router /false [get]
func (c *CommonController) False() {
	logs.Info(">>>> forward to false page start <<<<")
	c.TplName = "false.html"
}
