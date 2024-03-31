package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var dsn = "user:pass@tcp(117.51.148.112:3306)/study?charset=utf8"

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err = engine.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	defer engine.Close()

	Select()
	Insert()
	Update()
	Delete()
	Transaction()
}
