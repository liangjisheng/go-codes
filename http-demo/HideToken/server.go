package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析url传递的参数，默认不会解析，对于post则解析响应包的主体
	fmt.Println(r.Form) // 打印服务端
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	fmt.Println()
	fmt.Fprintf(w, "Hello alice") // 输出到客户端
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		// 增加隐藏字段token，通过md5(时间截)来获取唯一值
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证token的合法性，防止表单的重复提交
			fmt.Println(token)
		} else {
			// 不存在token报错
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
	fmt.Println()
}

func main() {
	http.HandleFunc("/", sayHelloName)       // 设置访问的路由
	http.HandleFunc("/login", login)         // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
