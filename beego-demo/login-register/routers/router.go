package routers

import (
	"go-demos/beego-demo/login-register/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, `get:PageLogin`)
	beego.Router("/register", &controllers.UserController{}, `post:Register`)
	beego.Router("/reallogin", &controllers.UserController{}, `post:Reallogin`)
}
