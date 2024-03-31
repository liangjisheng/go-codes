package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type responseToClient struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	var dataForm map[string][]string
	// 标识一个客户端的连接
	fmt.Println("client connect success ", r.RemoteAddr)
	// 获取地址栏内容
	fmt.Println(r.Method, r.RequestURI)
	// 获取请求头内容
	fmt.Println("request header:")
	for k, v := range r.Header {
		fmt.Println(k, v[0])
	}
	fmt.Println()
	data := make(map[string]string)
	if err := r.ParseForm(); err != nil {
		if r.Form != nil {
			dataForm = r.Form
		}
	}

	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Scheme:", r.URL.Scheme)

	// 读取客户端的内容
	buf := make([]byte, 2048)
	n, _ := r.Body.Read(buf)
	// 获取请求体中的内容
	fmt.Println("receive data from body", string(buf[:n]))

	if r.Method == "GET" {
		// r.ParseForm()
		for k, v := range r.Form {
			fmt.Println(k, ":", strings.Join(v, ""))
			data[k] = v[0]
		}
	}

	// 处理客户端发送的POST请求和PUT请求
	if r.Method == "POST" || r.Method == "PUT" {
		ct, ok := r.Header["Content-Type"]
		if ok {
			// 如果是json数据根据请求头判断
			if ct[0] == "application/json" {
				json.Unmarshal(buf[:n], &data)
			}
			// 如果是POST表单数据
			if ct[0] == "application/x-www-form-urlencoded" {
				if dataForm != nil {
					for k, v := range dataForm {
						data[k] = v[0]
					}
				}
			}
		}
	}

	// 处理客户端的DELETE请求
	if r.Method == "DELETE" {

	}

	// 记录当前时间 `2006-01-02 15:04:05` 是指的格式格式
	data["time"] = time.Now().Format("2006-01-02 15:04:05")
	m := responseToClient{200, "success", data}
	mjson, e := json.Marshal(m)
	if e != nil {
		fmt.Println(e)
	}
	// 以json格式响应给客户端
	fmt.Fprintf(w, "%v\n", string(mjson))
}

func defaultFunc1(w http.ResponseWriter, r *http.Request) {
	w.Write(nil)
}

// curl "http://127.0.0.1:8080/"
// curl "http://127.0.0.1:8080/?name=alice&toturial=gostudy&topic=web"

func main() {
	http.HandleFunc("/", defaultFunc1)
	addr := "127.0.0.1:8080"
	log.Println("http server start at:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
