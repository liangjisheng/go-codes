package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HelloWorldHandler ...
type HelloWorldHandler struct{}

func (handler *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK) // 设置响应状态码为 200
	fmt.Println("name:", string([]rune(params["name"])))
	fmt.Fprintf(w, "你好, %s!", params["name"]) // 发送响应到客户端
}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)                // 设置响应状态码为 200
	fmt.Fprintf(w, "Hello, %s", params["name"]) // 发送响应到客户端
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "文章列表")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "发布文章")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "修改文章")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "删除文章")
}

func showPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "文章详情")
}

func main() {
	r := mux.NewRouter()

	// curl "http://127.0.0.1:8080/hello"
	// curl "http://127.0.0.1:8080/hello/alice"
	// 正则表达式限制参数字符
	// 限定请求方法
	r.HandleFunc("/hello/{name:[a-z]+}", sayHelloWorld).Methods("GET", "POST")

	// windows 下会乱码,需要在浏览器上打开
	// curl "http://127.0.0.1:8080/zh/hello/你好"
	// curl "http://127.0.0.1:8080/zh/hello/alice"
	// curl -X POST "http://127.0.0.1:8080/zh/hello/alice"
	r.Handle("/zh/hello/{name}", &HelloWorldHandler{}).Methods("GET")

	// curl -X POST "http://127.0.0.1:8080/zh/hello1/alice"
	// 限定 Scheme
	r.Handle("/zh/hello1/{name}", &HelloWorldHandler{}).Methods("GET").Schemes("http")

	// curl "http://127.0.0.1:8080/request/header" -H "X-Requested-With: XMLHttpRequest"
	// 限定请求参数, 请求头必须包含 X-Requested-With 并且值为 XMLHttpRequest
	// 才可以访问指定路由 /request/header
	r.HandleFunc("/request/header", func(w http.ResponseWriter, r *http.Request) {
		header := "X-Requested-With"
		fmt.Fprintf(w, "包含指定请求头[%s=%s]", header, r.Header[header])
	}).Headers("X-Requested-With", "XMLHttpRequest")

	// curl "http://127.0.0.1:8080/query/string?token=test"
	// 通过 Queries 方法限定查询字符串, 查询字符串必须包含 token 且值为
	// test 才可以匹配到给定路由 /query/string
	r.HandleFunc("/query/string", func(w http.ResponseWriter, r *http.Request) {
		query := "token"
		fmt.Fprintf(w, "包含指定查询字符串[%s=%s]", query, r.FormValue(query))
	}).Queries("token", "test")

	// curl "http://127.0.0.1:8080/posts/"
	// curl -X POST "http://127.0.0.1:8080/posts/create"
	// curl -X PUT "http://127.0.0.1:8080/posts/update"
	// curl -X DELETE "http://127.0.0.1:8080/posts/delete"
	// curl "http://127.0.0.1:8080/posts/show"
	// 路由分组(基于子路由+路径前缀)
	// posts 前缀会应用到后面所有基于 postRouter 子路由定义的路由规则上
	postRouter := r.PathPrefix("/posts").Subrouter()
	// postRouter.HandleFunc("/", listPosts).Methods("GET")
	// postRouter.HandleFunc("/create", createPost).Methods("POST")
	// postRouter.HandleFunc("/update", updatePost).Methods("PUT")
	// postRouter.HandleFunc("/delete", deletePost).Methods("DELETE")
	// postRouter.HandleFunc("/show", showPost).Methods("GET")

	// 看一下 gorilla/mux 中的路由命名, 通过 Name 方法在路由规则中指定
	postRouter.HandleFunc("/", listPosts).Methods("GET").Name("posts.index")
	postRouter.HandleFunc("/create", createPost).Methods("POST").Name("posts.create")
	postRouter.HandleFunc("/update", updatePost).Methods("PUT").Name("posts.update")
	postRouter.HandleFunc("/delete", deletePost).Methods("DELETE").Name("posts.delete")
	postRouter.HandleFunc("/show", showPost).Methods("GET").Name("posts.show")

	// 打印路由对应的 URL
	indexURL, _ := r.Get("posts.index").URL()
	log.Println("文章列表链接:", indexURL)

	createURL, _ := r.Get("posts.create").URL()
	log.Println("发布文章链接:", createURL)

	showURL, _ := r.Get("posts.show").URL()
	log.Println("文章详情链接:", showURL)

	log.Println("http server start on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
