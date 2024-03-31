package leadybucket

import (
	"github.com/gin-gonic/gin"
	"math"
	"sync"
	"time"
)

//漏桶算法（Leaky Bucket），原理就是一个固定容量的漏桶，按照固定速率流出水滴
//漏桶限制的是常量流出速率（即流出速率是一个固定常量值），所以最大的速率就是出水的速率，不能应对出现突发流量

type LeakyBucket struct {
	rate float64 // 固定每秒出水速率
	capacity float64 // 桶的容量
	water float64 // 桶中当前水量
	lastLeakMs int64 // 桶上次漏水时间戳 ms
	lock sync.Mutex
}

func (l *LeakyBucket) Set(r, c float64) {
	l.rate = r
	l.capacity = c
	l.water = 0
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}

func (l *LeakyBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	eclipse := float64((now - l.lastLeakMs)) * l.rate / 1000 // 先执行漏水
	l.water = l.water - eclipse // 计算剩余水量
	l.water = math.Max(0, l.water) // 桶干了
	l.lastLeakMs = now

	if (l.water + 1) < l.capacity {
		// 尝试加水,并且水还未满
		l.water++
		return true
	} else {
		// 水满，拒绝加水
		return false
	}
}

// rate 桶流出速率, capacity 桶容量
func LimitMiddleware(rate, capacity int64) gin.HandlerFunc {
	var limiter LeakyBucket
	limiter.Set(float64(rate), float64(capacity))

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
