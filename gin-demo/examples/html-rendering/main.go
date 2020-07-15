package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// 使用 LoadHTMLGlob() 或者 LoadHTMLFiles()
	// r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	r.Run("127.0.0.1:8080")
}
