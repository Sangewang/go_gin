package model

type User struct {
	Username	string	 `json:"name" form:"name"`
	Age			uint8	 `json:"age"  form:"age"`
	Mobile		string	 `json:"mobile" form:"mobile"`
	Sex			string	 `json:"sex"  form:"sex"`
	Address		string	 `json:"address" form:"addreess"`
	Id			uint16	 `json:"id" form:"id"` 
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


// 但会升级后的用户信息
type UserInfo struct {
	Name	string	 `json:"name" form:"name"`
	Role	string	 `json:"role" form:"role"`
	Uid		string	 `json:"uid" form:"uid"`
	Avatar	string	 `json:"avatar"  form:"avatar"`
	Introduction string	 `json:"introduction" form:"introduction"` 
}