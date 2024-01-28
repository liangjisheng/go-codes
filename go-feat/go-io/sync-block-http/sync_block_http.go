package main

import (
	"fmt"
	"io"
	"net/http"
)

//这个示例中的IO操作都是同步阻塞的，因为我们使用了标准库中的函数来读取请求的Body和写入响应消息。
//这意味着当我们执行IO操作时，程序会一直等待直到操作完成，然后才会继续执行下一步操作。

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 读取请求的Body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// 处理请求
		response := processRequest(body)

		// 将响应写入ResponseWriter
		_, err = w.Write(response)
		if err != nil {
			http.Error(w, "Error writing response body", http.StatusInternalServerError)
			return
		}
	})

	// 启动HTTP服务器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}

func processRequest(request []byte) []byte {
	// 处理请求的逻辑
	return []byte("Processed request: " + string(request))
}
