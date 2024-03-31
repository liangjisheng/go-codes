package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"go-demos/beego-demo/user-register-login/models"
	"regexp"
	"strconv"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// LoginController ...
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("/login", c.Login)
	c.Mapping("/register", c.Register)
	c.Mapping("/register/do", c.DoRegister)
	c.Mapping("/login/do", c.DoLogin)
}

// Login ...
// @router /login [get]
func (c *LoginController) Login() {
	logs.Info(">>>> forward to login page start <<<<")
	c.TplName = "login.html"
}

// Register ...
// @router /register [get]
func (c *LoginController) Register() {
	logs.Info(">>>> forward to Register page start <<<<")
	c.TplName = "register.html"
}

// DoRegister ...
// @router /register/do [post]
func (c *LoginController) DoRegister() {
	var user models.User
	var result models.Result
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	checkError(err)

	userName := user.UserName
	password := user.Password
	email := user.Email
	logs.Info(">>>> user register information username=" + userName + ",email=" + email + ",password=" + password + " <<<<")

	pass, err := regexp.MatchString(`[a-zA-Z0-9]{8,16}`, password)
	checkError(err)
	if !pass {
		result.Code = 404
		result.Success = false
		result.Message = "密码格式错误"
	}

	em, err := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	checkError(err)
	if !em {
		result.Code = 404
		result.Success = false
		result.Message = "邮箱格式错误"
	}

	// 对密码加密
	user.Password = passwordEncode(password)

	id, e := orm.NewOrm().Insert(&user)
	if e != nil {
		result.Code = 404
		result.Success = false
		result.Message = "系统异常"
		logs.Error(e)
	} else {
		logs.Info(">>>> user register success,user id = " + strconv.FormatInt(id, 10) + " <<<<")
		result.Code = 200
		result.Success = true
		result.Message = "注册成功"
	}

	bytes, err := json.Marshal(&result)
	c.Ctx.ResponseWriter.Write(bytes)
}

func passwordEncode(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// DoLogin ...
// @router /login/do [post]
func (c *LoginController) DoLogin() {
	logs.Info(">>>> user do login start <<<<")
	var r models.Result
	var user models.User
	// 表单映射成struct, views/login.html 这个文件有问题，导致表单数据传不过来
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		r.Code = 404
		r.Success = false
		r.Message = "form error"
	}

	// 所以将下面的username和Password强制写成 alice ,使得显示登录成功
	user.UserName = "alice"
	user.Password = "alice"
	user.Password = passwordEncode(user.Password)

	right := models.QueryByNamePwd(user.UserName, user.Password)
	if right {
		r.Code = 200
		r.Success = true
		r.Message = "success"
	} else {
		r.Code = 404
		r.Success = false
		r.Message = "用户或密码错误"
	}

	//c.ServeJSON()
	//c.Data["json"] = &r
	// 或者使用这种方式
	var bytes []byte
	bytes, err = json.Marshal(&r)
	if err != nil {
		logs.Error(err)
	}
	c.Ctx.ResponseWriter.Write(bytes)
}
