package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		// 需要结束接收, 退出循环
		if err != nil {
			break
		}
	}
	// 函数已经结束, 发送通知
	wg.Done()
}

func main() {
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go socketRecv(conn, &wg)

	time.Sleep(time.Second)

	// 主动关闭套接字
	conn.Close()

	// 等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")
}
