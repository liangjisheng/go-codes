package handler

import (
	"go-demos/micro-demos/go-micro-demo/src/share/utils/log"

	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

// UserStatusHandler ...
type UserStatusHandler struct {
	pool          *redis.Pool
	logger        *zap.Logger
	namespace     string
	sessionExpire int
	tokenExpire   int
}

// NewUserStatusHandler ...
func NewUserStatusHandler(pool *redis.Pool) *UserStatusHandler {
	return &UserStatusHandler{
		pool:          pool,
		sessionExpire: 15 * 86400,
		tokenExpire:   15 * 86400,
		logger:        log.Instance().Named("UserStatusHandler"),
	}
}
