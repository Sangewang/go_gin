package model


// sys_dict_type表信息详情
// 返回数据的结构体 
type ResSysDictType struct {
	Did      int		`json:"dictId"`
	DictName string     `json:"dictName"`
	DictType string     `json:"dictType"`
	Status     string   `json:"status"`
	Remark	 string     `json:"remark"`
	CreateTime string   `json:"create_time"`
}

// 查询的结构体
type SysDictType struct {
	Did      int		`json:"dict_id" gorm:"auto-increment"`
	DictName string     `json:"dict_name"`
	DictType string     `json:"dict_type"`
	Status     string   `json:"status"`
	Remark	 string     `json:"remark"`
	CreateTime string   `json:"create_time"`
}

