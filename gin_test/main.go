package main

import (
	"fmt"
	db "gin_test/database"
)

func main () {
	fmt.Println("test go mod to download gin")
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}