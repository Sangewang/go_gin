package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func GetDataBase() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Baidu_1234"
	dbName := "gin_demo"
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database db = ", db)
	return db
}