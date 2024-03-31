package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "swagger/swag/docs"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @BasePath /api/v1
// @Schemes http
// @Host 127.0.0.1:8080
func main() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//gin-swagger 同时还提供了 DisablingWrapHandler 函数, 方便我们通过设置某些环境变量来禁用 swagger
	//r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.GET("/api/v1/hello/:id", hello)
	r.POST("/api/v1/helloPost", helloPost)
	r.POST("/api/v1/hello1", hello1)
	r.POST("/api/v1/hello2", hello2)

	r.Run("127.0.0.1:8080")
}

// ShowAccount  godoc
// @Summary     Get pet from the store
// @Description get string by ID
// @Accept      json
// @Produce     json
// @Param       id path string true "id"
// @Param       id2 query string true "id2"
// @Param       id3 query integer true "id3"
// @Param       id4 query number true "id4"
// @Param       id5 query boolean true "id5"
// @Success     200 {string} string	"ok"
// @Failure     404 {string} string "We need ID!!"
// @Router      /hello/{id} [get]
func hello(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		c.String(http.StatusOK, "hello %s", id)
		return
	}
	c.String(http.StatusOK, id)
}

// @Summary Add a new pet to the store post
// @Description insert string
// @Accept  json
// @Produce json
// @Success 200 {string} string "insert ok"
// @Failure 404 {string} string "insert fail"
// @Router /helloPost [post]
func helloPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "insert ok")
}

// @Summary Get pet from the store
// @Description Get pet from the store
// @Accept  json
// @Produce json
// @Param   ReqHello1 body object true "ReqHello1"
// @Param   uuid header string true "uuid"
// @Param   userID formData string true "userID"
// @Success 200 {object} ResHello1
// @Failure 404 {object} ResHello1
// @Failure 400 {object} ResponseError
// @Router /hello1 [post]
func hello1(ctx *gin.Context) {
	var req ReqHello1
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: "bad param",
		})
		return
	}

	uuid := ctx.GetHeader("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, ResponseError{
			Code:    400,
			Message: "bad param",
		})
		return
	}

	ctx.JSON(http.StatusOK, ResHello1{
		Name: "hello1",
	})
}

// @Summary pet from the store
// @Description Get pet from the store
// @Accept  x-www-form-urlencoded
// @Produce json
// @Param   userID formData string true "userID"
// @Success 200 {string} string "userID"
// @Failure 404 {string} string "fail"
// @Failure 400 {string} string "bad request"
// @Router /hello2 [post]
func hello2(ctx *gin.Context) {
	userID, ok := ctx.GetPostForm("userID")
	if !ok {
		ctx.JSON(http.StatusOK, "fail")
		return
	}

	ctx.JSON(http.StatusOK, userID)
}
