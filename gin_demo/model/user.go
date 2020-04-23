package model

type User struct {
	Username	string	 `json:"name" form:"name"`
	Age			uint8	 `json:"age"  form:"age"`
	Mobile		string	 `json:"mobile" form:"mobile"`
	Sex			string	 `json:"sex"  form:"sex"`
	Address		string	 `json:"address" form:"addreess"`
	Id			uint16	 `json:"id" form:"id"` 
}


type Login struct {
    User     string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
    Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}