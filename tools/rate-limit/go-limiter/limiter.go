package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

//LimitMiddleware 限流器中间件 每秒最高允许 maxRate 个请求
//最开始有 maxRate 个 token
func LimitMiddleware(maxRate int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxRate), maxRate)
	return func(ctx *gin.Context) {
		if !limiter.Allow() { // 这里用于控制当限流器达到阈值的判断
			ctx.JSON(http.StatusTooManyRequests, map[string]interface{}{
				"code": -1,
				"msg":  "too many requests",
				"data": nil,
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
