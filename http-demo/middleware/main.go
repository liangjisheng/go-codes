package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//创建两个中间件，一个用于记录程序执行的时长，另外一个用于验证请求用的是否是指定的 HTTPMethod，
//创建完后再用定义的 Chain函数把 http.HandlerFunc和应用在其上的中间件链起来，中间件会按添加顺序依次执行，最后执行到处理函数

type Middleware func(http.HandlerFunc) http.HandlerFunc

// 记录每个URL请求的执行时长
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 中间件的处理逻辑
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			// 调用下一个中间件或者最终的 handler 处理程序
			f(w, r)
		}
	}
}

// 验证请求用的是否是指定的HTTP Method，不是则返回 400 Bad Request
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

// 把应用到http.HandlerFunc处理器的中间件
// 按照先后顺序和处理器本身链起来供http.HandleFunc调用
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// 最终的处理请求的http.HandlerFunc
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	log.Println("listen on :8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}