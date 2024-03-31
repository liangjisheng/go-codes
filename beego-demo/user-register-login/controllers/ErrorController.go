package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// ErrorController ...
type ErrorController struct {
	beego.Controller
}

// URLMapping ...
func (c *ErrorController) URLMapping() {
	c.Mapping("Error404", c.Error404)
	c.Mapping("Error500", c.Error500)
}

// Error404 ...
func (c *ErrorController) Error404() {
	logs.Info(">>>> forward to 404 error page <<<<")
	c.Data["Content"] = "抱歉，查找的内容不存在"

	c.TplName = "error/404.html"
}

// Error500 ...
func (c *ErrorController) Error500() {
	logs.Info(">>>> forward to 500 error page <<<<")
	c.Data["Content"] = "抱歉，系统错误，努力解决中"

	c.TplName = "error/500.html"
}
