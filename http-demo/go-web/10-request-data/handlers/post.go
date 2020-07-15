package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// GetPosts ...
func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}

// Post ...
type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// AddPost ...
func AddPost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	// io.WriteString(w, string(body))
	post := Post{}
	json.Unmarshal(body, &post)
	fmt.Fprintf(w, "%#v\n", post)
}

// EditPost ...
func EditPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// request 对象上的 Form 可以获取所有请求参数, 包括查询字符串和请求实体
	// log.Fprintln(w, r.Form)

	id := r.Form.Get("id")
	log.Println("post id:", id)
	// 如果只想获取请求实体(即 POST 表单中的数据), 可以通过 PostForm 实现
	log.Println("form data:", r.PostForm)

	// 这两个方法的时候只能获取特定请求数据，不能一次获取所有请求数据
	log.Println("post id:", r.FormValue("id"))
	log.Println("post title:", r.PostFormValue("title"))
	log.Println("post title:", r.PostFormValue("content"))

	io.WriteString(w, "表单提交成功")
}
