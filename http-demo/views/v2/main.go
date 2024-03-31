package main

import (
	"html/template"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 存放用户数据
type UserData struct {
	Name string
	Text string
}

// 渲染页面并输出
func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, err := template.New(file).ParseFiles("views/" + file)
	checkErr(err)
	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}

// 处理主页请求
func index(w http.ResponseWriter, r *http.Request) {
	// 渲染页面并输出
	renderHTML(w, "index.html", "no data")
}

// 处理用户提交的数据
func page(w http.ResponseWriter, r *http.Request) {
	// 我们规定必须通过 POST 提交数据
	if r.Method == "POST" {
		// 解析客户端请求的信息
		if err := r.ParseForm(); err != nil {
			log.Println("Handler:page:ParseForm: ", err)
		}

		// 获取客户端输入的内容
		u := UserData{}
		u.Name = r.Form.Get("username")
		u.Text = r.Form.Get("usertext")

		// 渲染页面并输出
		renderHTML(w, "page.html", u)
	} else {
		// 如果不是通过 POST 提交的数据，则将页面重定向到主页
		renderHTML(w, "redirect.html", "/")
	}
}

func main() {
	http.HandleFunc("/", index)              // 设置访问的路由
	http.HandleFunc("/page", page)           // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
