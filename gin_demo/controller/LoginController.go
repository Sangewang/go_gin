package controller

import (
	"fmt"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
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


type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
}
 
var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

const(
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin = "请重新登陆"
)

func NewConfirmLoginEmail(context *gin.Context) {
	fmt.Println(">>>> NewConfirmLoginEmail action start <<<<")
	username := context.PostForm("email")
	password := context.PostForm("password")
	fmt.Println("username = ", username, "password = ", password)

	claims := &JWTClaims{
		Username:    username,
		Password:    password,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, _:=getToken(claims)

	// 后期加一个认证环节 => 返回用户名 密码和 token
	var s model.Login
	// sMap := make(map[int]model.Login)
	s.Username = username
	s.Password = password
	s.Token = signedToken
	// sMap[0] = s
	context.JSON(200, gin.H{
		"code" : 0,
		"msg"  : "验证成功",
		"data" : s,
	})

}

func getToken(claims *JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "",errors.New(ErrorReason_ServerBusy)
	}
	return signedToken,nil
}