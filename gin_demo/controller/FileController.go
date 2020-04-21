package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const BASE_NAME = "./static/file/"

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