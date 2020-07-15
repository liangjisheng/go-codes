package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		value, _ := c.Cookie("name")
		c.String(http.StatusOK, "Cookie:%s\n", value)
	})

	router.GET("/set", func(c *gin.Context) {
		c.SetCookie("name", "ljs", 10, "/", "localhost", false, true)
	})

	router.GET("/clc", func(c *gin.Context) {
		c.SetCookie("name", "ljs", -1, "/", "localhost", false, true)
	})

	router.Run("127.0.0.1:8080")
}
