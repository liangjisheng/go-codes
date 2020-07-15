package routes

import (
	"net/http"

	"uploadfile/handlers"
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
}
