package main

import (
	"fmt"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	store := persistence.NewInMemoryStore(time.Second)

	r.GET("/ping", func (c *gin.Context) {
		c.String(200, "pong " + fmt.Sprint(time.Now().Unix()))
	})

	// Cached Page
	r.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

	r.Run("127.0.0.1:8080")
}
