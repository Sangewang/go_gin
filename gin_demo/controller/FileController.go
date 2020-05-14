package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"reflect"
	"net/http"
	"encoding/json"
	"gin_demo/model"
)

const BASE_NAME = "../pic/"

func RenderView(context *gin.Context) {
	println(">>>> render to file upload view action start <<<<")
	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200, "fileUpload.html", gin.H{})
}

func FormUpload(context *gin.Context) {
	println(">>>> single upload file by form action start <<<<")
	fh, err := context.FormFile("file")
	println("fh = ", fh)
	checkError(err)
	fileName := fh.Filename
	println("fileName = ", fileName)

	file, err := fh.Open()
	defer file.Close()

	bytes, e := ioutil.ReadAll(file)
	e = ioutil.WriteFile(BASE_NAME + fileName, bytes, 0666)
	checkError(e)

	if e != nil {
		context.JSON(200,gin.H{
			"success": false,
		})
	} else {
		context.JSON(200,gin.H{
			"success": true,
		})
	}
}

func MultiUpload(context *gin.Context) {
	println(">>>> multi upload file by form action start <<<<")
	form, err := context.MultipartForm()
	checkError(err)
	files := form.File["files"]
	name := form.File["name"]
	email := form.File["email"]
	println("name = ", name, "email = ", email)

	var er error
	for _, f := range files {
		// 使用gin自带的保存文件方法
		println("f.Filename = ", f.Filename)
		er = context.SaveUploadedFile(f, BASE_NAME + f.Filename)
		checkError(err)
	}

	if er != nil {
		context.JSON(200, gin.H {
			"success" : false,
		})
	} else {
		context.JSON(200, gin.H{
			"success": true,
		})
	}
}

type TFile model.FileSys
var gFile TFile // 用一个全局变量获取File
type FileList []TFile

func (u TFile) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs u.Name = ", u.Name)
	gFile = u
}

func HandleUploadFile(context *gin.Context) {
	println(">>>> multi upload file by form action start <<<<")
	fmt.Println("context = ", context)
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

	// var datajson model.Login
	demoStr := string(data)
	// 转换成unit8类型的数组
	dat := []byte(demoStr)
	fmt.Println("dat = ", dat)
	// 定义一个 map
	var m map[string]FileList
    // 注意：反序列化 map，不需要 make，因为 make 操作被封装到了 Unmarsha 函数中
    err := json.Unmarshal(dat, &m)
    if err != nil {
        fmt.Println(err)
	}

	fmt.Println(" m = ", m)
	filelist := m["filelist"]
	fmt.Println("reflect(filelist) = ", reflect.TypeOf(filelist))

	t := reflect.TypeOf(filelist).Kind() //reflect.TypeOf(mm).Kind()获得变量类型，发现是slice
	v := reflect.ValueOf(filelist)  //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	fmt.Println("reflect.TypeOf(filelist).Kind = ", t)
	fmt.Println("reflect.ValueOf(filelist) = ", v)
	fmt.Println("type:", v.Type())

	for i := 0; i < v.Len(); i ++ {
		fmt.Println(v.Index(i))
		eachFile := v.Index(i)
		fmt.Println("eachFile type:", eachFile.Type())

		methodValue := eachFile.MethodByName("ReflectCallFuncNoArgs")
		args := make([]reflect.Value, 0)
		fmt.Println("args = ", args)
		methodValue.Call(args)
		fmt.Println("gFile.Name = ", gFile.Name, "gFile.Size = ", gFile.Size, "gFile.Url = ", gFile.Url)
		downloadReadFile(gFile.Name, gFile.Url)
	}
}

func downloadReadFile(filename string, targetUrl string) {
	resp, _ := http.Get(targetUrl)
	fmt.Println("resp = ", resp)
}