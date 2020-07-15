package controllers

import (
	"errors"
	"fmt"

	"demo.book.com/conf"
	"demo.book.com/models"
	"demo.book.com/services"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

// DemoController ...
type DemoController struct {
	Ctx iris.Context
}

// GetRecord1 /demo/record1
func (c *DemoController) GetRecord1() {
	driveSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", conf.SysConfMap["dbuser"], conf.SysConfMap["dbpwd"], conf.SysConfMap["dbhost"], conf.SysConfMap["dbport"], conf.SysConfMap["dbname"])
	engine, _ := xorm.NewEngine("mysql", driveSource)
	var info models.BookTb

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)

	engine.Table("book_tb").Where("id=?", 1).Get(&info)

	c.Ctx.JSON(info)
}

// GetOrm /demo/orm
func (c *DemoController) GetOrm() {
	service := services.NewBookService()
	// ID获取单条数据
	info := service.Get(1)
	// 获取列表
	list := service.GetList("Press = '湖南文艺出版社'", "ID asc", 2)
	// 获取分页
	total, pageList := service.GetPageList("", "ID asc", 0, 2)
	c.Ctx.JSON(
		iris.Map{
			"list":     list,
			"info":     info,
			"pageList": pageList,
			"total":    total,
		})
}

// GetXML /demo/x/m
func (c *DemoController) GetXML() {
	service := services.NewBookService()
	info := service.Get(1)
	c.Ctx.XML(info)
}

// GetErr 故意报错 /demo/err
func (c *DemoController) GetErr() {
	// 引发一个恐慌 程序会自动捕获并返回错误信息
	panic(errors.New("i'm a painc"))
}

// GetQPS /demo/q/p
func (c *DemoController) GetQPS() {
	c.Ctx.WriteString("hello")
}

// GetConf /demo/conf
func (c *DemoController) GetConf() {
	reload := c.Ctx.URLParam("reload")
	if reload != "" {
		// 如果有更新配置，重新读取配置文件
		conf.ReLoad()
	}
	c.Ctx.JSON(conf.SysConfMap)
}
