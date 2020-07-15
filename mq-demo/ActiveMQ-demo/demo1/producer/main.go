package main

import (
	"fmt"
	"time"

	"github.com/go-stomp/stomp"
)

func main() {
	conn, err := stomp.Dial("tcp", "117.51.148.112:61613")
	if err != nil {
		fmt.Println("dial err =", err)
		return
	}

	for i := 0; i < 10; i++ {
		// 调用conn下的send方法，接收三个参数
		// 参数一: 队列的名字
		// 参数二: 数据类型 一般是文本类型 直接写text/plain即可
		err := conn.Send("testQ", "text/plain", []byte(fmt.Sprintf("message:%d", i)))
		if err != nil {
			fmt.Println("send message err =", err)
		}
	}

	// 这里为什么要sleep一下 那就是conn.Send这个过程是不阻塞的
	// 相当于Send把数据放到了一个channel里面
	// 另一个goroutine从channel里面去取数据再放到消息队列里面
	// 但是还没等到另一个goroutine放入数据 此时循环已经结束了
	// 因此最好要sleep一下 根据测试 如果不sleep 那么发送1000条数据
	// 最终进入队列的大概是980条数据 这说明了什么
	// 说明了当程序把1000条数据放到channel里面的时候，另一个goroutine只往队列里面放了980条
	// 剩余的20条还没有来得及放入 程序就结束了
	time.Sleep(time.Second * 1)
}
