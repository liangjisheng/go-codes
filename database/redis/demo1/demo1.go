package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

// github上一个人fork了github.com/garyburd/redigo/redis驱动，然后更新的一些bug
func demo1() {
	var client goredis.Client
	// 端口为redis默认端口
	client.Addr = "127.0.0.1:6379"

	// 字符串操作
	key := "foo"
	client.Set(key, []byte("bar"))
	val, _ := client.Get(key)
	fmt.Println(string(val))
	client.Del(key)

	// list操作
	key = "l1"
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush(key, []byte(v))
	}
	dbvals, _ := client.Lrange(key, 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del(key)
}
