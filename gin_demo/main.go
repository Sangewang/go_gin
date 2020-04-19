package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin_demo/controller"
)

// https://github.com/ypcfly/helloGin
func main() {
	// Engin
	router := gin.Default()
	// 加载template下的模板
	router.LoadHTMLGlob("template/*")
	router.GET("/hello", hello) // hello 函数处理"/hello"的请求

	// 路由组
	user := router.Group("/user")
	{	// 请求参数在请求路由上
		user.GET("/get/:id/:username", controller.QueryById)
		user.GET("/query", controller.QueryParam)
		user.POST("/insert", controller.InsertNewUser)
		user.GET("/form", controller.RenderForm)
		user.POST("/form/post", controller.PostForm)
		// 可以添加其它的路由请求

	}

	// 指定端口号和地址
	router.Run(":8080")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")
	context.JSON(http.StatusOK, gin.H {
		"code": 200,
		"success": true,
	})
}