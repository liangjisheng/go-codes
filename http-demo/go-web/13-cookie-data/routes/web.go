package routes

import (
	"net/http"

	"cookiedata/handlers"
)

// WebRoute 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// WebRoutes 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	// curl "http://127.0.0.1:8080/"
	WebRoute{
		"Home",
		"GET",
		"/",
		handlers.Home,
	},
	// curl "http://127.0.0.1:8080/posts"
	WebRoute{
		"Posts",
		"GET",
		"/posts",
		handlers.GetPosts,
	},
	// curl "http://127.0.0.1:8080/user/1"
	WebRoute{
		"User",
		"GET",
		"/user/{id}",
		handlers.GetUser,
	},
	// curl -X "POST" "http://127.0.0.1:8080/post/add" -id '{"title":"test", "content":"hello"}' -H "Content-Type: application/json"
	WebRoute{
		"NewPost",
		"POST",
		"/post/add",
		handlers.AddPost,
	},
	// curl -X "POST" "http://127.0.0.1:8080/post/edit?id=1" -id "title=test&content=hello"
	WebRoute{
		"UpdatePost",
		"POST",
		"/post/edit",
		handlers.EditPost,
	},
	// curl -i "http://127.0.0.1:8080/post/edit?id=1&title=test&content=hello"
	WebRoute{
		"UpdatePost",
		"GET",
		"/post/edit",
		handlers.EditPost,
	},
	WebRoute{
		"UploadImage",
		"POST",
		"/image/upload",
		handlers.UploadImage,
	},
	WebRoute{
		"UploadImage1",
		"POST",
		"/image/upload1",
		handlers.UploadImage1,
	},
	// curl -i "http://127.0.0.1:8080/error"
	WebRoute{
		"ApiError",
		"GET",
		"/error",
		handlers.Error,
	},
	// 如果是在浏览器中访问的话, 页面就会跳转到重定向的页面
	// curl -i "http://127.0.0.1:8080/redirect"
	WebRoute{
		"Redirect",
		"GET",
		"/redirect",
		handlers.Redirect,
	},
	// 浏览器访问
	// http://localhost:8080/setcookies
	WebRoute{
		"SetCookie",
		"GET",
		"/setcookies",
		handlers.SetCookie,
	},
	WebRoute{
		"GetCookie",
		"GET",
		"/getcookies",
		handlers.GetCookie,
	},
	// 在浏览器中访问 http://localhost:8080/set_welcome_message, 页面会重定向到
	// http://localhost:8080/get_welcome_message, 然后打印出欢迎消息, 说明 Cookie 读取成功
	WebRoute{
		"SetMessage",
		"GET",
		"/set_welcome_message",
		handlers.SetWelcomeMessage,
	},
	WebRoute{
		"GetMessage",
		"GET",
		"/get_welcome_message",
		handlers.GetWelcomeMessage,
	},
}
