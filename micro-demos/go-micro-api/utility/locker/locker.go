package locker

import (
	"errors"
	"time"

	"demo/utility/db"

	"github.com/gomodule/redigo/redis"
)

// Locker ...
type Locker struct {
	Key   string
	Error error
}

// Lock ...
func Lock(key string) (locker *Locker) {
	locker = &Locker{Key: key}

	conn := db.GetRedis()
	defer conn.Close()

	// EX 表示过期时间
	// 利用 Redis set key 时的一个 NX 参数可以保证在这个 key 不存在的情况下写入成功
	// 并且再加上 EX 参数可以让该 key 在超时之后自动删除
	r, _ := redis.String(conn.Do("SET", key, 1, "EX", 10, "NX"))
	if r != "OK" {
		locker.Error = errors.New("locker failed")
	}
	return
}

// TryLock ...
func TryLock(key string, timeout time.Duration) (locker *Locker) {
	locker = &Locker{Key: key}

	conn := db.GetRedis()
	defer conn.Close()

	start := time.Now()
	for time.Now().Sub(start) < timeout {
		reply, _ := redis.String(conn.Do("SET", key, 1, "EX", 60, "NX"))
		if reply == "OK" {
			return
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}

	locker.Error = errors.New("locker timeout")
	return
}

// Close ...
func (lock *Locker) Close() {
	if lock.Error == nil {
		conn := db.GetRedis()
		defer conn.Close()

		conn.Do("DEL", lock.Key)
	}
}
