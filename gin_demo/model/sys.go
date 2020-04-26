package model


// sys_dict_type表信息详情
type TSysDictType struct {
	Did      int		`json:"dictId"`
	Dictname string     `json:"dictName"`
	Dicttype string     `json:"dictType"`
	Status     string   `json:"status"`
	Remark	 string     `json:"remark"`
	CreateTime string   `json:"create_time"`
}