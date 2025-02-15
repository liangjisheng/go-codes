package controller

import (
	"go-demos/projects/cgo/cgo"
	"go-demos/projects/cgo/constant"
	"go-demos/projects/cgo/service"
	"go-demos/projects/cgo/utils"
	"log"
	"net/http"
)

/**
 * r.PostFormValue  : 可以解析 Post/PUT Content-Type=application/x-www-form-urlencoded 或 Content-Type=multipart/form-data
 */

type UserController struct {
	cgo.ApiController
}

var userService = new(service.UserService)

func (p *UserController) Router(router *cgo.RouterHandler) {
	router.Router("/register", p.register)
	router.Router("/login", p.login)
	router.Router("/findAll", p.findAll)
	router.Router("/findUser", p.findUser)
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserController) register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		cgo.ResultFail(w, "username or password can not be empty")
		return
	}

	id := userService.Insert(username, password)
	if id <= 0 {
		cgo.ResultFail(w, "register fail")
		return
	}
	cgo.ResultOk(w, "register success")
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *UserController) login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if utils.Empty(username) || utils.Empty(password) {
		cgo.ResultFail(w, "username or password can not be empty")
		return
	}

	users := userService.SelectUserByName(username)
	if len(users) == 0 {
		cgo.ResultFail(w, "user does not exist")
		return
	}
	if users[0].Password != password {
		cgo.ResultFail(w, "password error")
		return
	}

	// session
	session := cgo.GlobalSession().SessionStart(w, r)
	session.Set(constant.KEY_USER, &users[0])
	cgo.ResultOk(w, "login success")
}

// GET/POST
func (p *UserController) findAll(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("GSESSION")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(cookie.Value)
	}

	users := userService.SelectAllUser()
	cgo.ResultJsonOk(w, users)
}

func (p *UserController) findUser(w http.ResponseWriter, r *http.Request) {
	user := p.GetUser(w, r)
	cgo.ResultJsonOk(w, user)
}
