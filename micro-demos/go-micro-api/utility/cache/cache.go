package cache

import (
	"errors"
	"time"

	"demo/utility/db"

	"github.com/gomodule/redigo/redis"
)

// Set ...
func Set(key, val string, ttl time.Duration) error {
	conn := db.GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("SET", key, val, "EX", ttl.Seconds()))
	if err != nil {
		return err
	}
	if r != "OK" {
		return errors.New("NOT OK")
	}
	return nil
}

// Get ...
func Get(key string) (string, error) {
	conn := db.GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return r, nil
}

// Del ...
func Del(key string) (int, error) {
	conn := db.GetRedis()
	defer conn.Close()

	return redis.Int(conn.Do("DEL", key))
}
