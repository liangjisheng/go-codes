package tokenbucket

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

//有一个固定的桶，桶里存放着令牌（token）。一开始桶是空的，系统按固定的时间（rate）往桶里添加令牌，
//直到桶里的令牌数满，多余的请求会被丢弃。当请求来的时候，从桶里移除一个令牌，如果桶是空的则拒绝请求或者阻塞

type TokenBucket struct {
	rate         int64 // 固定的token放入速率, r/s
	capacity     int64 // 桶的容量
	tokens       int64 // 桶中当前token数量
	lastTokenSec int64 // 桶上次放token的时间戳 s

	lock sync.Mutex
}

func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.tokens = 0
	l.lastTokenSec = time.Now().Unix()
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().Unix()
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 先添加令牌
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now

	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

// LimitMiddleware rate 固定的token放入速率 r/s, capacity 桶容量
func LimitMiddleware(rate, capacity int64) gin.HandlerFunc {
	var limiter TokenBucket
	limiter.Set(rate, capacity)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			//log.Print("request not allow")
			c.Abort()
			return
		}

		//log.Print("request allow")
		c.Next()
	}
}
