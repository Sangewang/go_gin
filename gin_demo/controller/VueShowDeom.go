package controller

import (
	"github.com/gin-gonic/gin"
)

func ShowVueDemo(context *gin.Context) {
	println(">>>> render to show vue view action start <<<<")
	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200, "vueDemo.html", gin.H{})
}