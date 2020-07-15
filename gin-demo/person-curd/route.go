package main

import (
	"go-demos/gin-demo/person-curd/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.Index)
	router.POST("/person", apis.AddPerson)
	router.GET("/persons", apis.GetPersons)
	router.GET("/person/:id", apis.GetPerson)
	router.PUT("/person/:id", apis.UpdatePerson)
	router.DELETE("/person/:id", apis.DelPerson)

	return router
}
