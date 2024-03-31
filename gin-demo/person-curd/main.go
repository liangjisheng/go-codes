package main

import (
	db "go-demos/gin-demo/person-curd/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}
