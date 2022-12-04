package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Payload struct {

}

func (p *Payload) UpdateToS3() error {
	//存储逻辑,模拟操作耗时
	time.Sleep(500 * time.Millisecond)
	fmt.Println("update ok.")
	return nil
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	// 业务过滤
	// 请求body解析......
	var p Payload
	go p.UpdateToS3()
	w.Write([]byte("操作成功\n"))
}

// curl "http://127.0.0.1:8080/payload"

func main() {
	http.HandleFunc("/payload", payloadHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
