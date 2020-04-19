package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"net/http"
	"gin_demo/database"
	"gin_demo/model"
)

// 定义数据库全局变量
var db *sql.DB

// 初始化函数先执行
func init () {
	log.Println(">>>> Get Database Connection start <<<<")
	db = database.GetDataBase()
}

// 查询参数 localhost:8080/user/query?id=2&name=hello 
// http://127.0.0.1:8080/user/query?id=1&name=林旺
func QueryParam(context *gin.Context) {
	println(">>>> query user by url params action start <<<<")
	id := context.Query("id")
	name := context.Request.URL.Query().Get("name")
	
	rows, err := db.Query("select username, address, age, mobile, sex from t_user where id = ? and username = ?", id, name)
	checkError(err)

	var username string
	var address string
	var age int 
	var mobile string
	var sex string

	for rows.Next() {
        err := rows.Scan(&username, &address, &age, &mobile, &sex)
        checkError(err)
		fmt.Println("username = ", username, "address = ", address, "age = ", age, "mobile = ", mobile, "sex = ", sex)
	}

	context.JSON(200, gin.H{
		"用户名": username,
		"地址": address,
		"年龄": age,
		"电话": mobile,
		"性别": sex,
	})
}

// 查询信息 localhost:8080/user/get/1/林旺
func QueryById (context *gin.Context) {
	println(">>>> get user by id and name action start <<<<")

	// 获取请求参数
	id := context.Param("id")
	name := context.Param("username")

	println("id = ", id, "username = ", name)

	//查询数据
    rows, err := db.Query("select username, address, age, mobile, sex from t_user where id = ? and username = ?", id, name)
	checkError(err)
	
	var username string
	var address string
	var age int 
	var mobile string
	var sex string

    for rows.Next() {
        err := rows.Scan(&username, &address, &age, &mobile, &sex)
		checkError(err)
		fmt.Println("username = ", username, "address = ", address, "age = ", age, "mobile = ", mobile, "sex = ", sex)
	}

	context.JSON(200, gin.H{
		"用户名": username,
		"地址": address,
		"年龄": age,
		"电话": mobile,
		"性别": sex,
	})
}

// 插入信息 
func InsertNewUser (context *gin.Context) {
	println(">>>> insert controller action start <<<< => ", context.PostForm)

	name := context.PostForm("name")
	age := context.PostForm("age")
	address := context.PostForm("address")
	mobile := context.PostForm("mobile")
	sex := context.PostForm("sex")
	println("name = ", name, "age = ", age, "address = ", address, "mobile = ", mobile, "sex = ", sex)

	var user model.User

	// 直接将结构体和提交的json参数作绑定
	err := context.ShouldBindJSON(&user)

	// mysql 插入数据
    stmt, err := db.Prepare("INSERT t_user SET username=?,sex=?,address=?,mobile=?,age=?")
    checkError(err)

    res, err := stmt.Exec(name, sex, address, mobile, age)
	checkError(err)

	var count int64
	count, err = res.RowsAffected()
	checkError(err)

	if count != 1 {
		context.JSON(200, gin.H{
			"success": false,
		})
	} else {
		context.JSON(200, gin.H{
			"success": true,
		})
	}
}

// form 提交表单
func PostForm(context *gin.Context) {
	println(">>>> bind form post params action start <<<< => ", context.PostForm)
	name := context.PostForm("name")
	age := context.PostForm("age")
	address := context.PostForm("address")
	mobile := context.PostForm("mobile")
	sex := context.PostForm("sex")
	println("name = ", name, "age = ", age, "address = ", address, "mobile = ", mobile, "sex = ", sex)

	var u model.User
	println(" u = ", &u)
	// 绑定参数到结构体
	context.ShouldBindJSON(&u)
	// context.ShouldBind(&u)
	println("u.Username = ", u.Username)

	// mysql 插入数据
    stmt, err := db.Prepare("INSERT t_user SET username=?,sex=?,address=?,mobile=?,age=?")
    checkError(err)

    res, err := stmt.Exec(name, sex, address, mobile, age)
	checkError(err)
	
	// 查询影响的行数
	var count int64
	count, err = res.RowsAffected()
	checkError(err)
	
	if count != 1 {
		context.JSON(200, gin.H{
			"success": false,
		})
	} else {
		// 重定向
		context.Redirect(http.StatusMovedPermanently, "/file/view")
	}
}

// 跳转html
func RenderForm(context *gin.Context) {
	println(">>>> render to html action start <<<<")
	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200, "insertUser.html", gin.H{})
}

// 报错信息
func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}