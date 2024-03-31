package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

var cnt int64
var key = "alice"
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			lock(func() {
				cnt++
				fmt.Printf("after incr is %d\n", cnt)
			}, i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("cnt = %d\n", cnt)
}

func lock(handler func(), id int) {
	defer wg.Done()

	ctx := context.TODO()
	lockSuccess, err := redisClient.SetNX(ctx, key, 1, time.Second*3).Result()
	if err != nil || lockSuccess != true {
		fmt.Println(id, " get lock fail", err)
		return
	} else {
		fmt.Println(id, " get lock success")
	}

	handler()

	//unlock
	_, err = redisClient.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(id, "unlock fail", err)
	} else {
		fmt.Println(id, "unlock success")
	}
}

//上述版本会出现一个问题：当某个goroutine1执行时间比较长，例如操作一个10GB的大文件
//goroutine2去获取锁是发现goroutine1虽然有锁但是过期了，goroutine2就毫不客气的拿到了该锁，
//然后goroutine2去执行业务代码，goroutine2也执行了很久。 go调度器在goroutine2 执行期间，
//goroutine1调取执行，这时候goroutine1并不知道自己因为超时而时去了该锁，而对该锁进行了删除
//这时goroutine3 去抢占锁成功了，就会出现goroutine2和goroutine3同时操作互斥资源的情况
//那么怎么解决该问题呢? 每个goroutine 对锁设置不同的标签做为值，每个goroutine在删除锁
//之前读取一下锁的值，确保是自己持有的情况下，才会进行删除锁的操作

func lockV2(handler func()) {
	//lock
	ctx := context.TODO()
	uuid := getUUID()
	lockSuccess, err := redisClient.SetNX(ctx, key, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return
	} else {
		fmt.Println("get lock success")
	}

	//run func
	handler()

	//unlock
	value, _ := redisClient.Get(ctx, key).Result()
	if value == uuid { //compare value,if equal then del
		_, err := redisClient.Del(ctx, key).Result()
		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock success")
		}
	}
}

func getUUID() string {
	return uuid.NewV4().String()
}
