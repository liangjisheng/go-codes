package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // 解析参数
	fmt.Println(r.Form)             // 在服务端打印请求参数
	fmt.Println("URL:", r.URL.Path) // 请求 URL
	fmt.Println("Scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Println()
	fmt.Fprintf(w, "hello alice") // 发送响应到客户端
}

// curl "http://127.0.0.1:8080/"
// curl "http://127.0.0.1:8080/?name=alice&toturial=gostudy&topic=web"

func main() {
	http.HandleFunc("/", sayHelloWorld)
	log.Println("http server start on 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
