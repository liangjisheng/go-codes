package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

// BeforeExecFilter ...
func BeforeExecFilter(ctx *context.Context) {
	logs.Info(">>>> BeforeExec filter start <<<<")
}

// BeforeRouterFilter ...
func BeforeRouterFilter(ctx *context.Context) {
	logs.Info(">>>> BeforeRouter filter start <<<<")
}

// BeforeStaticFilter 第一个执行
func BeforeStaticFilter(ctx *context.Context) {
	logs.Info(">>>> BeforeStatic filter start <<<<")
	// session只能在之后的filter使用
	//id := ctx.Input.Session("id").(string)
	//println(id)
}

// AfterExecFilter ...
func AfterExecFilter(ctx *context.Context) {
	logs.Info(">>>> AfterExec filter start <<<<")
}

// FinishRouterFilter ...
func FinishRouterFilter(ctx *context.Context) {
	logs.Info(">>>> FinishRouter filter start <<<<")
}
