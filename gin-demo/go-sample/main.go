package main

import (
	"go-demos/gin-demo/go-sample/models"
	"go-demos/gin-demo/go-sample/router"
)

func init() {
	models.Setup()
}

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r := router.SetupRouter()
	r.Run("127.0.0.1:8080") // 在 127.0.0.1:8080 上监听并服务
}
