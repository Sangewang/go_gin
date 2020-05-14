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
	Did      int		`gorm:"column:dict_id"`  
	DictName string     
	DictType string    
	Status     string  
	Remark	 string     
	CreateTime string   
}


/*
// 查询的结构体
type SysDictType struct {
	Did      int		`json:"dict_id"`
	DictName string     `json:"dict_name"`
	DictType string     `json:"dict_type"`
	Status     string   `json:"status"`
	Remark	 string     `json:"remark"`
	CreateTime string   `json:"create_time"`
}


type Animal struct {
    AnimalId    int64     `gorm:"column:beast_id"`         // 设置列名为`beast_id`
    Birthday    time.Time `gorm:"column:day_of_the_beast"` // 设置列名为`day_of_the_beast`
    Age         int64     `gorm:"column:age_of_the_beast"` // 设置列名为`age_of_the_beast`
}*/