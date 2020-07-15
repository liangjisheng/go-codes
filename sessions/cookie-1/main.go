package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// curl http://127.0.0.1:8080/incr
	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	router.Run("127.0.0.1:8080")
}
