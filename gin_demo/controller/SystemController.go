package controller

import (
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gin_demo/model"
	"gin_demo/database"
	"github.com/jinzhu/gorm"
)

var gormdb *gorm.DB
// 初始化函数先执行
func init () {
	log.Println(">>>> Get Sys Database Connection start <<<<")
	gormdb = database.GetSysTableDb()
	gormdb.SingularTable(true)
	gormdb.LogMode(true)
}

func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func GetSysDictTypeInfo(context *gin.Context) {
	log.Println(">>>> func GetSysDictTypeInfo start <<<<")
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	//SysData := make([]model.SysDictType, 0)
	newdb := gormdb

	var SysData []model.SysDictType

	fmt.Println("SysData = ", SysData)
	demoStr := string(data)
	// 转换成unit8类型的数组
	dat := []byte(demoStr)
	fmt.Println("dat = ", dat)
	// 定义一个 map
    var m map[string]interface{}
    // 注意：反序列化 map，不需要 make，因为 make 操作被封装到了 Unmarsha 函数中
    err := json.Unmarshal(dat, &m)
    if err != nil {
        fmt.Println(err)
	}

	dictName := m["qdictName"]
	dictType := m["qdictType"]
	status := m["qstatus"]

	if dictName != "All" {     
		newdb = newdb.Where("dict_name = ?", dictName)
	}
		
	if dictType != "All" {  
		newdb = newdb.Where("dict_type = ?", dictType)
	}
		
	if status != "All" { 
		statusDict := map[interface {}]int{ "正常":0, "异常":1}
		fmt.Println("typeof(status) = ", typeof(status))
		fmt.Println("typeof(statusDict) = ", typeof(statusDict))
		tmpStatus := statusDict[status]
		fmt.Println("tmpStatus = ", tmpStatus)
		// statusDict := make(map[string]int)
		// statusDict["正常"] = 0
		// statusDict["异常"] = 1
		newdb = newdb.Where("status = ?", tmpStatus)
	}
	fmt.Println("newdb = ", newdb)
	if err := newdb.Select("dict_id,dict_name,dict_type,status,remark,create_time").Find(&SysData).Error; err != nil {
        fmt.Println(err.Error())
	}
	dataLen := len(SysData)
	flag := false
	fmt.Println("SysData = ", SysData, "len(SysData) = ", len(SysData))
	// 结构体数组
	fmt.Println("typeof(SysData) = ", typeof(SysData))

	var s model.ResSysDictType
	var gCount = 0
	sMap := make(map[int]model.ResSysDictType)

	if dataLen > 0 {
		flag = true
		// 单个结构体
		fmt.Println("SysData[0] = ", SysData[0])
		fmt.Println("typeof(SysData[0]) = ", typeof(SysData[0]))
		// 查询结构体每个变量
		fmt.Println("SysData[0].dict_id = ", SysData[0].Did)
		// fmt.Println("SysData[0][0] = ", SysData[0])

		for i :=0 ; i < dataLen; i ++ {
			eachData :=  SysData[i]
			s.Did = eachData.Did
			s.DictName = eachData.DictName
			s.DictType = eachData.DictType
			s.Status = eachData.Status
			s.Remark = eachData.Remark
			s.CreateTime = eachData.CreateTime
			sMap[gCount] = s
			gCount += 1
		}
	}

	if flag != false {
		context.JSON(200, gin.H{
			"code" : -1,
			"msg"  : "验证失败",
			"data" : sMap,
			"old_data": SysData,
		})
	} else {
		// 输出序列化后的结果
		// fmt.Printf("序列化后 = %v\n", string(SysData))
		context.JSON(200, gin.H{
			"code" : 0,
			"msg"  : "验证成功",
			"data"  : sMap,
			"old_data": SysData,
		})
	}

}
/*
func GetSysDictTypeInfo(context *gin.Context) {
	log.Println(">>>> func GetSysDictTypeInfo start <<<<")
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	// var datajson model.Login
	demoStr := string(data)
	// 转换成unit8类型的数组
	dat := []byte(demoStr)
	fmt.Println("dat = ", dat)
	// 定义一个 map
    var m map[string]interface{}
    // 注意：反序列化 map，不需要 make，因为 make 操作被封装到了 Unmarsha 函数中
    err := json.Unmarshal(dat, &m)
    if err != nil {
        fmt.Println(err)
	}
	dictName = m["qdictName"]
	dictType = m["qdictType"]
	status = m["qstatus"]

	fmt.Println("basicSql = ", basicSql)
	rows, err := db.Query("select dict_id, dict_name, dict_type, status, remark, create_time from sys_dict_type")
	fmt.Println("rows = ", rows, "err = ", err)
	checkError(err)

	var s model.TSysDictType
	var gCount = 0
	sMap := make(map[int]model.TSysDictType)
	for rows.Next() {
        var uid int
        var dictname string
		var dicttype string
		var status string
		var remark string
		var create_time string
		err = rows.Scan(&uid, &dictname, &dicttype, &status, &remark, &create_time)
		s.Did = uid
		s.Dictname = dictname
		s.Dicttype = dicttype
		s.Status = status
		s.Remark = remark
		s.CreateTime = create_time
		sMap[gCount] = s
		gCount += 1
        checkError(err)
	}
	fmt.Println("结构体字典 sMap = ", sMap)
	postData, err := json.Marshal(&sMap)
	if err != nil{
		context.JSON(200, gin.H{
			"code" : -1,
			"msg"  : "验证失败",
			"data" : postData,
		})

	} else {
		//输出序列化后的结果
		fmt.Printf("序列化后 = %v\n", string(postData))
		context.JSON(200, gin.H{
			"code" : 0,
			"msg"  : "验证成功",
			"data"  : sMap,
			"json_data" : postData,
		})
	}
}
*/