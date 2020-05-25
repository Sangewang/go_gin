package model

// 用户登录的结构体
type Login struct {
    Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	Token string `json:"token"`
}

