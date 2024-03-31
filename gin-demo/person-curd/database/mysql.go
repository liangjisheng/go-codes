package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SqlDB ...
var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/05blog?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
