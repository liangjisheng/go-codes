package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTP响应body的读取影响TCP连接的复用，如果想要复用TCP连接以提高传输速率，就需要读取response的body内容

func getStatus(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	//通过 wireshark 抓包可知
	//当我们没有读取 response.Body 的内容时，每次 HTTP GET 请求，都会进行三次握手，创建自己的TCP连接

	//当读取 response.Body 时, 发现这次只进行了一次三次握手
	io.Copy(io.Discard, resp.Body)

	return resp.StatusCode, nil
}

func main() {
	url := "https://www.baidu.com"
	// 连续进行两次Get请求
	for i := 0; i < 2; i++ {
		status, err := getStatus(url)
		fmt.Printf("status:%d, err(%+v)\n", status, err)
		time.Sleep(1 * time.Second)
	}
}
