package model


const BASE_NAME = "../pic/"


type PicSys struct {
	Name string `gorm:"column:pic_name"`
	// Raw  map `json:"raw"`
	Size int `gorm:"column:pic_size"`
	Percentage int `json:"percentage"`
	Status string `json:"status"`
	Uid int `json:"uid"`
	Url string `json:"url"`
 }


 type RetPicSys struct {
	Name string `json:"name"`
	Url []string `json:"url"`
 }

