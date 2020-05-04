package model

type FileSys struct {
	Name string `json: "name"`
	Size int `json: "size"`
	Percentage int `json:"percentage"`
	Status string `json:"status"`
	Uid int `json:"uid"`
	Url string `json:"url"`
 }