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

// curl -i -v "http://127.0.0.1:8080/hello"
// curl "http://127.0.0.1:8080/hello/alice"
// windows 下会乱码,需要在浏览器上打开
// curl "http://127.0.0.1:8080/zh/hello/你好"
// curl "http://127.0.0.1:8080/zh/hello/alice"

func main() {
	r := mux.NewRouter()
	// 正则表达式限制参数字符
	r.HandleFunc("/hello/{name:[a-z]+}", sayHelloWorld)
	r.Handle("/zh/hello/{name}", &HelloWorldHandler{})

	log.Println("http server start on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
