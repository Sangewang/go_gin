package model

type User struct {
	Username	string	 `json:"name" form:"name"`
	Age			uint8	 `json:"age"  form:"age"`
	Mobile		string	 `json:"mobile" form:"mobile"`
	Sex			string	 `json:"sex"  form:"sex"`
	Address		string	 `json:"address" form:"addreess"`
	Id			uint16	 `json:"id" form:"id"` 
}

// 用户登录的结构体
type Login struct {
    User     string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`	
}

// TUser表信息详情
type TUserDetailInfo struct {
	Uid      int		`json:"uid"`
	Username string     `json:"username"`
	Province string     `json:"province"`
	City     string     `json:"city"`
	Address  string     `json:"address"`
	Code	 string     `json:"code"`
	Detail	 string     `json:"detail"`
	CreateTime string   `json:"create_time"`
}