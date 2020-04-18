package main 

import (
	"github.com/gin-gonic/gin"
	. "gin_test/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.POST("/person", AddPersonApi)
	router.GET("/persons", GetPersonsApi)
	router.GET("/person/:id", GetPersonsApi)
	router.PUT("/person/:id", ModPersonApi)
	router.DELETE("/person/:id", DelPersonApi)
	return router
}