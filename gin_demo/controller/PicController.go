package controller

import (
	"fmt"
	"log"
	"os"
	"io"
	"io/ioutil"
	"net/http"
	// "reflect"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"gin_demo/model"
	// "gin_demo/database"
	// "github.com/jinzhu/gorm"
)



// type来临时定义一个和model.Pic具有同样功能的类型
type TPic model.PicSys

type PicList []TPic


var ret TPic // 用一个全局变量获取Pic
func (u TPic) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs u.Name = ", u.Name)
	ret = u
}

// 图片上传处理

func New_PicUploadHandle(context *gin.Context) {
	log.Println(">>>> func PicUploadHandle start <<<<")
	file, header, err := context.Request.FormFile("file0") 
	if err != nil {
		// context.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		fmt.Println(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	
	fmt.Println("file = ", file, "filename = ", filename)

	out, err := os.Create(model.BASE_NAME + filename)
	if err != nil {
	  log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
	  log.Fatal(err)
	}
}

func PicUploadHandle(context *gin.Context) {
	println(">>>> multi upload file by form action start <<<<")
	form, err := context.MultipartForm()
	checkError(err)
	files := form.File["files"]

	var er error
	for _, f := range files {
		// 使用gin自带的保存文件方法
		println("f.Filename = ", f.Filename)
		er = context.SaveUploadedFile(f, model.BASE_NAME + f.Filename)
		checkError(err)
	}
	fmt.Println("er = ", er)
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
