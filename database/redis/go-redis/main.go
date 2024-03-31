package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// 定义一组常量
const (
	redisIP   = "117.51.148.112"
	redisPort = "6379"
	redisPwd  = "password"
	redisDB   = 0
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":" + redisPort, // ip:port
		Password: redisPwd,                  // redis连接密码
		DB:       redisDB,                   // 选择的redis库
		PoolSize: 20,                        // 设置连接数,默认是10个连接
	})
}

func main() {
	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	base()
	// strings()
	// hash()
	// list()
	// set()
	// zset()
	// subscribe()
}
