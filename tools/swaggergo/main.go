package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "go-demos/tools/swaggergo/docs"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
func main() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/v1/hello/:id", hello)
	r.Run("127.0.0.1:8080")
}

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce json
// @Param   id path string true "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 404 {string} string "We need ID!!"
// @Router /hello/{id} [get]
func hello(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		c.String(http.StatusOK, "hello %s", id)
		return
	}
	c.String(http.StatusForbidden, "We need ID!!")
}
