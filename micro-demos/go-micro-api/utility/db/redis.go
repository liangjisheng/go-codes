package db

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

// InitRedis 初始化redis
func InitRedis(addr string, db, maxIdle, maxOpen int) (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxOpen,
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialDatabase(db))
		},
	}

	conn := GetRedis()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed")
	}
	return
}

// GetRedis 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// CloseRedis 关闭redis
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}
