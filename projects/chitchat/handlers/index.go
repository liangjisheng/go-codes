package handlers

import (
	"net/http"

	"chitchat/models"
)

// Index 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	// 主布局文件 layout.html
	// 顶部导航模板 navbar.html
	// 首页视图模板 index.html
	// 编译多个视图模板时, 默认以第一个模板名作为最终视图模板名

	threads, err := models.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	}

	// files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	// templates := template.Must(template.ParseFiles(files...))
	// threads, err := models.Threads()
	// if err == nil {
	// 	templates.ExecuteTemplate(w, "layout", threads)
	// }
}

// Err ...
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}
