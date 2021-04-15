package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 支持子目录路径
	// http.StripPrefix() 方法配合 http.Handle() 或
	// http.HandleFunc() 可以实现带路由前缀的文件服务
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("./tmp"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

	// http://localhost:8080/tmpfiles/
}
