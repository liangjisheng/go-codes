package main

import (
	"fmt"
	"log"
	"net/http"
)

// 向客户端写入这些数据，以便客户端可以填写文本并提交
var indexHTML = `<html>
<head>
	<meta http-equiv="Content-type" content="text/html; charset=utf-8">
	<title>测试</title>
</head>
<body>
	<form action="/page" method="post">
		用户名：<br>
		<input name="username" type="text"><br>
		请输入文本：<br>
		<textarea name="usertext"></textarea><br>
		<input type="submit" value="提交">
	</form>
</body>
</html>`

// 用于将页面重定向到主页
var redirectHTML = `<html>
<head>
	<meta http-equiv="Content-type" content="text/html; charset=utf-8">
	<meta http-equiv="Refresh" content="0; url={{.}}">
</head>
<body></body>
</html>`

// 处理主页请求
func index(w http.ResponseWriter, r *http.Request) {
	// 向客户端写入我们准备好的页面
	fmt.Fprintf(w, indexHTML)
}

// 处理客户端提交的数据
func page(w http.ResponseWriter, r *http.Request) {
	// 我们规定必须通过 POST 提交数据
	if r.Method == "POST" {
		// 解析客户端请求的信息
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		// 获取客户端输入的内容
		userName := r.Form.Get("username")
		userText := r.Form.Get("usertext")
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["usertext"])
		// 将内容反馈给客户端
		fmt.Fprintf(w, "你好 %s，你输入的内容是：%s", userName, userText)
	} else {
		// 如果不是通过 POST 提交的数据，则将页面重定向到主页
		fmt.Fprintf(w, redirectHTML)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/page", page)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
