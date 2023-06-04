package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"sync"
)

func NewIPLimiter(rps int) gin.HandlerFunc {
	limiters := &sync.Map{}

	return func(c *gin.Context) {
		// 获取限速器
		// key 除了 ip 之外也可以是其他的，例如 header，user name 等
		key := c.ClientIP()
		l, _ := limiters.LoadOrStore(key, ratelimit.New(rps))
		now := l.(ratelimit.Limiter).Take()
		fmt.Printf("now: %s\n", now)
		c.Next()
	}
}
