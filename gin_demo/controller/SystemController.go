package controller

import (
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"gin_demo/model"
)

func GetSysDictTypeInfo(context *gin.Context) {
	log.Println(">>>> func GetSysDictTypeInfo start <<<<")
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	
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