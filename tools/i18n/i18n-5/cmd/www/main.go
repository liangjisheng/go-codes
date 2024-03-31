package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	// 初始化一个路由实例，并为主页添加路径和处理程序
	mux := pat.New()
	//mux.Get("/:locale", http.HandlerFunc(handleHome))
	mux.Get("/:locale", http.HandlerFunc(handleHomeLocalizer))

	// 使用路由实例启动HTTP服务器。
	log.Println("starting server on :8008...")
	err := http.ListenAndServe(":8008", mux)
	log.Fatal(err)
}