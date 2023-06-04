package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	e := gin.Default()
	// 新建一个限速器，允许突发 10 个并发，限速 3rps，超过 500ms 就不再等待
	e.Use(NewIPLimiter(3, 10, 500*time.Millisecond))
	e.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	e.Run(":18080")
}
