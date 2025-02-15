package controllers

import (
	"demo.book.com/conf"
	"demo.book.com/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// BookController ...
type BookController struct {
	Ctx iris.Context
}

// Get /book
func (c *BookController) Get() mvc.Result {
	service := services.NewBookService()
	list := service.GetList("", "ID asc", 0)
	return mvc.View{
		Name: "book/home.html",
		Data: iris.Map{
			"Title":  "首页-" + conf.SysConfMap["port"],
			"List":   list,
			"Server": conf.SysConfMap["port"],
		},
		Layout: "shared/bookLayout.html",
	}
}

// GetAjaxbooks /book/ajaxbooks?key=go 访问地址是小写的
func (c *BookController) GetAjaxbooks() {
	key := c.Ctx.URLParam("key")
	service := services.NewBookService()
	list := service.GetList(" bookName like '%"+key+"%'", "ID asc", 0)
	c.Ctx.JSON(list)
}
