package controllers

import (
	"encoding/json"
	"go-demos/beego-demo/user-register-login/models"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// UserController ...
type UserController struct {
	beego.Controller
}

// URLMapping 配置注解路由
func (u *UserController) URLMapping() {
	u.Mapping("QueryById", u.QueryByID)
	u.Mapping("QueryList", u.QueryList)
	u.Mapping("Register", u.Register)
	u.Mapping("SaveUser", u.SaveUser)
	u.Mapping("UserJSON", u.UserJSON)
}

// Register ...
// @router /user [get]
func (u *UserController) Register() {
	logs.Info(">>>> forward to register page start <<<<")
	u.TplName = "user/form.html"
}

// QueryByID ...
// @router /user/get/:id [get]
func (u *UserController) QueryByID() {
	strID := u.Ctx.Input.Param(":id")
	logs.Info(">>>> query user by userId start <<<<")
	id, err := strconv.Atoi(strID)
	user := models.QueryUserByID(id)
	checkError(err)

	u.Data["User"] = user
	bytes, err := json.Marshal(user)
	u.Ctx.ResponseWriter.Write(bytes)

	//u.TplName = "user/user.html"
}

// QueryList ...
// @router /user/list [get]
func (u *UserController) QueryList() {
	logs.Info(">>>> query user list start <<<<")

	// 数据库查询所有用户
	users := models.QueryUserList()

	//bytes,err := json.Marshal(users)
	//checkError(err)
	//u.Ctx.ResponseWriter.Write(bytes)

	u.Data["List"] = users
	u.TplName = "user/list.html"
}

// SaveUser ...
// @router /user/save [post]
func (u *UserController) SaveUser() {
	logs.Info(">>>> save register information start <<<<")
	// 获取表单数据
	//form := u.Ctx.Request.PostForm
	//username := form.Get("username")
	//age := form.Get("age")
	//sex := form.Get("sex")
	//mobile := form.Get("mobile")

	// 表单转struct
	var user models.User
	err := u.ParseForm(&user)
	checkError(err)
	// 校验...

	// 写入数据库，返回id值
	id := models.InsertUser(&user)
	println(id)
	u.Ctx.Redirect(302, "/success")
}

// UserJSON ...
// @router /json/save [post]
func (u *UserController) UserJSON() {
	logs.Info(">>>> save user json information start <<<<")
	// requestBody数据
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	checkError(err)

	// 数据库处理
	id := models.InsertUser(&user)
	println("insert user id=" + strconv.FormatInt(id, 10))

	u.Ctx.Redirect(302, "/success")
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
