package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//此方法弊端是对超时时间的设置有要求，需要根据具体业务设置一个合理的经验值，避免锁超时时间到了，业务没执行完的问题

func main() {
	rds, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer rds.Close()

	for true {
		// 检查是否有锁与加锁必须是原子性操作
		result, err := rds.Do("SET", "lock", 1, "EX", 5, "NX")
		if err != nil {
			fmt.Println("redis set error.", err)
			return
		}
		result, err = redis.String(result, err)
		// 加锁失败，继续轮询
		if result != "OK" {
			fmt.Println("SET lock failed.")
			time.Sleep(5 * time.Second)
			continue
		}

		// 加锁成功
		fmt.Println("work begin")
		// 此处处理业务
		time.Sleep(5 * time.Second)
		fmt.Println("work end")

		// 业务处理结束后释放锁
		result, err = rds.Do("del", "lock")
		if err != nil {
			fmt.Println("redis del err", err)
		}
		break
	}
}
