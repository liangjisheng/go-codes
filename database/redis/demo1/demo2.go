package main

// redigo 驱动演示如何进行数据操作

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := "127.0.0.1:6379"
	Pool = newPool(redisHost)
	close()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

// Get get k-v
func Get(key string) ([]byte, error) {
	fun := "Get"

	conn := Pool.Get()
	defer conn.Close()

	if err := Pool.TestOnBorrow(conn, time.Now()); err != nil {
		fmt.Printf("%s err: %v\n", fun, err)
		return []byte{}, err
	}

	var data []byte
	data, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func demo2() {
	test, err := Get("test")
	if err != nil {
		fmt.Printf("main err: %v", err)
		return
	}

	fmt.Println(string(test))
}
