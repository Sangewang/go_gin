package controller

import (
	"fmt"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
)

func ConfirmLoginEmail(context *gin.Context) {
	fmt.Println(">>>> ConfirmLoginEmail action start <<<<")
	username := context.PostForm("email")
	password := context.PostForm("password")
	fmt.Println("username = ", username, "password = ", password)

	// 后期加一个认证环节 => 返回用户名 密码和 token
	var s model.Login
	// sMap := make(map[int]model.Login)
	s.Username = username
	s.Password = password
	s.Token = "admin"
	// sMap[0] = s
	context.JSON(200, gin.H{
		"code" : 0,
		"msg"  : "验证成功",
		"data" : s,
	})
}

func ConfirmLoginOut(context *gin.Context) {
	fmt.Println(">>>> ConfirmLoginOut action start <<<<")
	context.JSON(200, gin.H{
		"code" : 0,
		"msg"  : "验证成功",
		"data" : "",
	})
}