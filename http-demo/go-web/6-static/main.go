package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 需要在 main.go 目录下执行
// go run main.go

func main() {
	r := mux.NewRouter()

	dir := "static"
	// curl "http://127.0.0.1:8080/static/app.js"
	// curl "http://127.0.0.1:8080/static/test.jpg"
	// 处理形如 http://localhost:8080/static/<filename> 的静态资源路由
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	log.Println("http server start on 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
