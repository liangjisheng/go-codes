package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"time"
)

//Redis采用Lua脚本可以执行多个个性化的原子操作

const (
	SCRIPT_LOCK = ` 
    local res=redis.call('GET', KEYS[1])
    if res then
        return 0
    else
        redis.call('SET',KEYS[1],ARGV[1]);
        redis.call('EXPIRE',KEYS[1],ARGV[2])
        return 1
    end 
    `

	SCRIPT_EXPIRE = ` 
    local res=redis.call('GET', KEYS[1])
    if not res then
        return -1
    end 
    if res==ARGV[1] then
        redis.call('EXPIRE', KEYS[1], ARGV[2])
        return 1
    else
        return 0
    end 
    `

	SCRIPT_DEL = ` 
    local res=redis.call('GET', KEYS[1])
    if not res then 
        return -1
    end 
    if res==ARGV[1] then
        redis.call('DEL', KEYS[1])
    else
        return 0
    end 
    `
)

var (
	UUID = getUUID()
)

func ResetExpire() {
	fmt.Println("Reset expire begin...")
	rds, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis server error.", err)
		return
	}

	for true {
		luaExpire := redis.NewScript(1, SCRIPT_EXPIRE)
		result, err := redis.Int(luaExpire.Do(rds, "lock", UUID, 5))
		if err != nil {
			fmt.Println("luaExpire exec error", err)
			break
		}
		if result != 1 {
			fmt.Println("Reset expire failed.")
			break
		} else {
			fmt.Println("Reset expire succeed.")
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("Reset expire end.")
}

func getUUID() string {
	return uuid.NewV4().String()
}

func main() {
	rds, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis server error.", err)
		return
	}
	defer rds.Close()

	luaLock := redis.NewScript(1, SCRIPT_LOCK)
	luaDel := redis.NewScript(1, SCRIPT_DEL)

	for true {
		result, err := redis.Int(luaLock.Do(rds, "lock", UUID, 5))
		if err != nil {
			fmt.Println("luaLock exec error.", err)
			time.Sleep(5 * time.Second)
			continue
		}
		if result == 0 {
			fmt.Println("Set lock failed.")
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Println("Set lock succeed.")

		go ResetExpire()
		// 加锁成功
		fmt.Println("work begin")
		// 此处处理业务
		time.Sleep(10 * time.Second)
		fmt.Println("work end")

		// 业务处理结束后释放锁
		result, err = redis.Int(luaDel.Do(rds, "lock"))
		if err != nil {
			fmt.Println("del lock failed err", err)
			fmt.Println("result", result)
		}
		return
	}
}
