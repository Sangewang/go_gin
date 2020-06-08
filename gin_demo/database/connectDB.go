package database

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// http://gorm.book.jasperxu.com/models.html#md 

var dbDriver = "mysql"
var dbUser = "root"
// var dbPass = "Baidu_1234"
var dbPass = ""
var dbName = "gin_demo"
var dbPort = 3306
var dbHost = "127.0.0.1"

func GetDataBase() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database db = ", db)
	return db
}

func GetSysTableDb() (db *gorm.DB) {
	var newdb *gorm.DB

	mysqlStr := "%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	conStr := fmt.Sprintf(mysqlStr, dbUser, dbPass, dbHost, dbPort, dbName)
	newdb, err := gorm.Open(dbDriver, conStr)
	if err != nil {
		panic("数据库连接失败")
	}
	return newdb
}
