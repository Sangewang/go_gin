package main

import (
	"fmt"
	// "time"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"gin_demo/controller"
	// "github.com/gin-contrib/cors"
)

// https://github.com/ypcfly/helloGin
func main() {
	// Engin
	router := gin.Default()

	// 允许使用跨域请求  全局中间件
	router.Use(Cors())

	// 加载template下的模板
	router.LoadHTMLGlob("template/*")
	router.GET("/hello", hello) // hello 函数处理"/hello"的请求

	// 路由组
	user := router.Group("/user")
	{	// 请求参数在请求路由上
		user.POST("/login", controller.ConfirmLoginInfo)
		//user.POST("/detail", controller.GetAllDetailInfo)
		user.GET("/get/:id/:username", controller.QueryById)
		user.GET("/query", controller.QueryParam)
		user.POST("/insert", controller.InsertNewUser)
		user.GET("/form", controller.RenderForm)
		user.POST("/form/post", controller.PostForm)
		// 可以添加其它的路由请求
	}

	file := router.Group("/file")
	{
		// 跳转到文件上传界面
		file.GET("/view", controller.RenderView)
		// 根据表单上传
		file.POST("/singleUpload", controller.FormUpload)
		file.POST("/multiUpload", controller.MultiUpload)
		// base64上传图片
		// file.POST("/upload", controller.Base64Upload)
	}

	vue := router.Group("/vue")
	{
		// vue 前端框架初探
		vue.GET("/demo", controller.ShowVueDemo)
	}

	// 指定端口号和地址
	router.Run(":9090")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")
	context.JSON(http.StatusOK, gin.H {
		"code": 200,
		"success": true,
	})
}


func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		fmt.Println("method = ", method)
        origin := c.Request.Header.Get("Origin")        //请求头部
        var headerKeys []string                             // 声明请求头keys
        for k, _ := range c.Request.Header {
            headerKeys = append(headerKeys, k)
        }
        headerStr := strings.Join(headerKeys, ", ")
        if headerStr != "" {
            headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
        } else {
            headerStr = "access-control-allow-origin, access-control-allow-headers"
        }
        if origin != "" {
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
            c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
            //  header的类型
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
            //              允许跨域设置                                                                                                      可以返回其他子段
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
            c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
            c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
            c.Set("content-type", "application/json")       // 设置返回格式是json
        }

        //放行所有OPTIONS方法
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "Options Request!")
        }
        // 处理请求
        c.Next()        //  处理请求
    }
}