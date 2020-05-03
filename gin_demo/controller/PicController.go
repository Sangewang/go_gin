package controller

import (
	"fmt"
	"log"
	"io/ioutil"
	"reflect"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gin_demo/model"
	// "gin_demo/database"
	// "github.com/jinzhu/gorm"
)

// type来临时定义一个和model.Pic具有同样功能的类型
type TPic model.PicSys

type PicList []TPic

/*
var gormdb *gorm.DB
// 初始化函数先执行
func init () {
	log.Println(">>>> Get Pic Database Connection start <<<<")
	gormdb = database.GetSysTableDb()
	gormdb.SingularTable(true)
	gormdb.LogMode(true)
}
*/

var ret TPic // 用一个全局变量获取Pic
func (u TPic) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs u.Name = ", u.Name)
	ret = u
}

// 图片上传处理
func PicUploadHandle(context *gin.Context) {
	log.Println(">>>> func PicUploadHandle start <<<<")

	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	fmt.Println("reflect(data) = ", reflect.TypeOf(string(data)))
	// var datajson model.Login
	demoStr := string(data)
	// 转换成unit8类型的数组
	dat := []byte(demoStr)
	fmt.Println("dat = ", dat)
	// 定义一个 map
	var m map[string]PicList
    // 注意：反序列化 map，不需要 make，因为 make 操作被封装到了 Unmarsha 函数中
    err := json.Unmarshal(dat, &m)
    if err != nil {
        fmt.Println(err)
	}

	fmt.Println(" m = ", m)
	picfile := m["picfile"]
	fmt.Println("reflect(picfile) = ", reflect.TypeOf(picfile))
	// fmt.Println(" picfile[0] = ", picfile[0])
	t := reflect.TypeOf(picfile).Kind() //reflect.TypeOf(mm).Kind()获得变量类型，发现是slice
	v := reflect.ValueOf(picfile)  //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	fmt.Println("reflect.TypeOf(picfile).Kind = ", t)
	fmt.Println("reflect.ValueOf(picfile) = ", v)
	fmt.Println("type:", v.Type())
	
	for i := 0; i < v.Len(); i++ {
		fmt.Println(v.Index(i))
		eachPic := v.Index(i)
		fmt.Println("eachPic type:", eachPic.Type())
		// 极限恶心方法 => 应该有更好的方案
		methodValue := eachPic.MethodByName("ReflectCallFuncNoArgs")
		args := make([]reflect.Value, 0)
		fmt.Println("args = ", args)
		methodValue.Call(args)
		fmt.Println("ret.Name = ", ret.Name, "ret.Size = ", ret.Size)
		// 准备插入数据库
		picsys := model.PicSys{Name: ret.Name, Size: ret.Size, Url: ret.Url}
		gormdb.NewRecord(picsys) // => 主键为空返回`true`
		gormdb.Create(&picsys)
	}
}	

// 图片获取函数
func PicDownloadHandle(context *gin.Context) {
	log.Println(">>>> func PicDownloadHandle start <<<<")

	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

	var PicData []model.PicSys
	
	gormdb.Find(&PicData)
	fmt.Println("PicData = ", PicData)

	var s model.RetPicSys
	var gCount = 0
	sMap := make(map[int]model.RetPicSys)

	// 目前只需要url和name，且url徐亚婆list类型
	for i :=0 ; i < len(PicData); i ++ {
		eachData :=  PicData[i]
		s.Name = eachData.Name
		s.Url = []string{eachData.Url}
		sMap[gCount] = s
		gCount += 1
	}
	flag := true
	if flag == false {
		context.JSON(200, gin.H{
			"code" : -1,
			"msg"  : "验证失败",
			"data" : sMap,
			"old_data": PicData,
		})
	} else {
		// 输出序列化后的结果
		// fmt.Printf("序列化后 = %v\n", string(SysData))
		context.JSON(200, gin.H{
			"code" : 0,
			"msg"  : "验证成功",
			"data"  : sMap,
			"old_data": PicData,
		})
	}
}
