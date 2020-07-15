package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:CommonController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:CommonController"],
        beego.ControllerComments{
            Method: "False",
            Router: `/false`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:CommonController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:CommonController"],
        beego.ControllerComments{
            Method: "Success",
            Router: `/success`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:FileController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:FileController"],
        beego.ControllerComments{
            Method: "Base64",
            Router: `/file/base64`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:FileController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:FileController"],
        beego.ControllerComments{
            Method: "DoUpload",
            Router: `/file/doUpload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:FileController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:FileController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/file/upload`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"],
        beego.ControllerComments{
            Method: "DoLogin",
            Router: `/login/do`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:LoginController"],
        beego.ControllerComments{
            Method: "DoRegister",
            Router: `/register/do`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserJSON",
            Router: `/json/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"],
        beego.ControllerComments{
            Method: "QueryByID",
            Router: `/user/get/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"],
        beego.ControllerComments{
            Method: "QueryList",
            Router: `/user/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"] = append(beego.GlobalControllerRouter["go-demos/beego-demo/user-register-login/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveUser",
            Router: `/user/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
