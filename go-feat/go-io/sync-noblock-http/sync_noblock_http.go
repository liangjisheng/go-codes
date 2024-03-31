package main

import (
	"fmt"
	"io"
	"net/http"
)

//我们使用了goroutine和channel来实现非阻塞IO操作。当一个HTTP请求到达时，我们创建了一个channel
//来接收请求的Body。然后，我们启动一个goroutine来执行读取请求Body的操作，并将读取到的Body发送到channel中

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建一个channel来接收请求的Body
		bodyChan := make(chan []byte)

		// 将请求的Body读取操作放到goroutine中执行
		go func() {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				// 发送错误信息到channel
				bodyChan <- []byte(fmt.Sprintf("Error reading request body: %v", err))
			} else {
				// 发送读取到的Body到channel
				bodyChan <- body
			}
		}()

		// 处理请求
		response := processRequest(bodyChan)

		// 将响应写入ResponseWriter
		_, err := w.Write(response)
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

func processRequest(bodyChan chan []byte) []byte {
	// 非阻塞地从channel中获取请求的Body
	select {
	case body := <-bodyChan:
		// 处理请求的逻辑
		return []byte("Processed request: " + string(body))
	default:
		// 如果没有收到请求的Body，则返回错误信息
		return []byte("No request body received")
	}
}
