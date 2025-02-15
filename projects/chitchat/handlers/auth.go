package handlers

import (
	"net/http"

	"chitchat/models"
)

// Login 登录页面
// GET /login
func Login(writer http.ResponseWriter, request *http.Request) {
	// t := parseTemplateFiles("auth.layout", "navbar", "login")
	// t.Execute(writer, nil)
	generateHTML(writer, nil, "auth.layout", "navbar", "login")
}

// Signup 注册页面
// GET /signup
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// SignupAccount 注册新用户
// POST /signup
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		// log.Println("Cannot parse form")
		danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		// log.Println("Cannot create user")
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// Authenticate 通过邮箱和密码字段对用户进行认证
// POST /authenticate
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		// log.Println("Cannot find user")
		danger(err, "Cannot find user")
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			// log.Println("Cannot create session")
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// Logout 用户退出
// GET /logout
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		// log.Println("Failed to get cookie")
		warning(err, "Failed to get cookie")
		session := models.Session{UUID: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
