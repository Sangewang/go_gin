package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	//"os"
	"encoding/json"
	"io/ioutil"
	"reflect"
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
    rows, err := db.Query("select username, address, age, mobile, sex, passwd from t_user where id = ? and username = ?", id, name)
	checkError(err)
	
	var username string
	var address string
	var age int 
	var mobile string
	var sex string
	var passwd string

    for rows.Next() {
        err := rows.Scan(&username, &address, &age, &mobile, &sex, &passwd)
		checkError(err)
		fmt.Println("username = ", username, "address = ", address, "age = ", age, "mobile = ", mobile, "sex = ", sex, "passwd = ", passwd)
	}

	context.JSON(200, gin.H{
		"用户名": username,
		"地址": address,
		"年龄": age,
		"电话": mobile,
		"性别": sex,
		"密码": passwd,
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
	passwd := context.PostForm("passwd")
	println("name = ", name, "age = ", age, "address = ", address, "mobile = ", mobile, "sex = ", sex, "passwd = ", passwd)

	var user model.User

	// 直接将结构体和提交的json参数作绑定
	err := context.ShouldBindJSON(&user)

	// mysql 插入数据
    stmt, err := db.Prepare("INSERT t_user SET username=?,sex=?,address=?,mobile=?,age=?,passwd=?")
    checkError(err)

    res, err := stmt.Exec(name, sex, address, mobile, age, passwd)
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
	passwd := context.PostForm("passwd")
	println("name = ", name, "age = ", age, "address = ", address, "mobile = ", mobile, "sex = ", sex, "passwd = ", passwd)

	var u model.User
	println(" u = ", &u)
	// 绑定参数到结构体
	context.ShouldBindJSON(&u)
	// context.ShouldBind(&u)
	println("u.Username = ", u.Username)

	// mysql 插入数据
    stmt, err := db.Prepare("INSERT t_user SET username=?,sex=?,address=?,mobile=?,age=?,passwd=?")
    checkError(err)

    res, err := stmt.Exec(name, sex, address, mobile, age, passwd)
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

// 用户登录认证
func ConfirmLoginInfo(context *gin.Context) {
	println(">>>> confirm user info is valid, receive json data <<<<")
	
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	fmt.Println("reflect(data) = ", reflect.TypeOf(string(data)))
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
	username := m["username"]
	password := m["password"]
	fmt.Println("反序列化结果：", m, "username = ", username, "password = ", password)

	// 验证用户名和密码是否正确

	rows, err := db.Query("select username, address, age, mobile, sex, passwd from t_user where username = ? and passwd = ?", username, password)
	fmt.Println("rows = ", rows, "err = ", err)
	checkError(err)

	var row_id int
	err = db.QueryRow("select id from t_user where username = ? and passwd = ?", username, password).Scan(&row_id)
	if err == sql.ErrNoRows {
		context.JSON(200, gin.H{
			"code" : -1,
			"msg"  : "验证失败",
			"data" : "nil",
			"jsondata" : "nil",
		})
	} else {
		log.Print("能查询到记录 = ", row_id)
		oldData := GetAllBasicInfo()
		jsonData, _ := json.Marshal(&oldData)
		context.JSON(200, gin.H{
			"code" : 0,
			"msg"  : "验证成功",
			"data" : oldData,
			"jsondata" : jsonData,
		})
	}

	/*
	//其实就是将request中的Body中的数据按照JSON格式解析到json变量中
	var datajson model.Login
	if err := context.ShouldBind(&datajson); err != nil {
		fmt.Println("error = ", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	*/
}

// 获取全部的细节信息
func GetAllBasicInfo() map[int]model.TUserDetailInfo {
	/*
	println(">>>> confirm GetAllDetailInfo, receive json data <<<<")
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	*/
	rows, err := db.Query("select id, username, province, city, address, code, detail, create_time from t_user_detail")
	fmt.Println("rows = ", rows, "err = ", err)
	checkError(err)

	var s model.TUserDetailInfo
	var gCount = 0
	sMap := make(map[int]model.TUserDetailInfo)
	for rows.Next() {
        var uid int
        var username string
		var province string
		var city string
		var address string
		var code string
		var detail string
		var create_time string
		err = rows.Scan(&uid, &username, &province, &city, &address, &code, &detail, &create_time)
		s.Uid = uid
		s.Username = username
		s.Province = province
		s.City = city
		s.Address = address
		s.Code = code
		s.Detail = detail
		s.CreateTime = create_time
		sMap[gCount] = s
		gCount += 1
        checkError(err)
	}
	fmt.Println("结构体字典 sMap = ", sMap)
	return sMap
	/*
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
			"old_data"  : sMap,
			"json_data" : postData,
		})
	}
	*/
}


// 获取全部的细节信息
func GetAllDetailInfo(context *gin.Context) {
	println(">>>> confirm GetAllDetailInfo, receive json data <<<<")
	data, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))
	
	rows, err := db.Query("select id, username, province, city, address, code, detail, create_time from t_user_detail")
	fmt.Println("rows = ", rows, "err = ", err)
	checkError(err)

	var s model.TUserDetailInfo
	var gCount = 0
	sMap := make(map[int]model.TUserDetailInfo)
	for rows.Next() {
        var uid int
        var username string
		var province string
		var city string
		var address string
		var code string
		var detail string
		var create_time string
		err = rows.Scan(&uid, &username, &province, &city, &address, &code, &detail, &create_time)
		s.Uid = uid
		s.Username = username
		s.Province = province
		s.City = city
		s.Address = address
		s.Code = code
		s.Detail = detail
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

