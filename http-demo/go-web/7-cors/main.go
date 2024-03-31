package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// curl -v "http://127.0.0.1:8080/api/cors"
	// 注意: 为了让中间件可以设置 CORS 头, 在 Methods 方法中必须包含 OPTIONS 方法
	r.HandleFunc("/api/cors", corsHandler).Methods(http.MethodGet, http.MethodPut, http.MethodOptions)
	// CORSMethodMiddleware 中间件会将上一步设置的方法设置到 Access-Control-Allow-Methods 响应头
	r.Use(mux.CORSMethodMiddleware(r))

	log.Println("http server start on 127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", r)
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("Cors Request"))
}
