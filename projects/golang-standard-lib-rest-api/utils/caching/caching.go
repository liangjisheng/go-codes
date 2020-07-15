package caching

import (
	"time"

	"github.com/go-redis/redis"
)

// Cache 我们使用了一个接口来抽象。这样的话，如果你决定使用另外的缓存服务时
// 代码修改就非常方便直接了
type Cache interface {
	Get(key string) (string, error)
	Set(key, value string, expiration time.Duration) error
}

// Redis use redis implement cache
type Redis struct {
	Client *redis.Client
}

// Get get the key
func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

// Set set the key
func (r *Redis) Set(key, value string, expiration time.Duration) error {
	return r.Client.Set(key, value, expiration).Err()
}

// Connect get redis client
func Connect(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
