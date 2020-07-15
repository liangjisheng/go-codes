package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Test ...
type Test struct {
	ID  int
	Num int
}

func createTables(db *gorm.DB) {
	db.CreateTable(&Test{})
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:ljs199711@tcp(localhost:3306)/study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open db error", err)
	} else {
		fmt.Println("open db success")
	}
}

func main() {
	if db != nil {
		defer db.Close()
	}

	if !db.HasTable("tests") {
		createTables(db)
	}

	test := Test{Num: 123456}
	db.Create(&test)
	fmt.Println("test.id is", test.ID)

	var tests []Test
	db.Find(&tests)
	fmt.Println(tests)

	for index, line := range tests {
		fmt.Println("index", index, " line ", line)
	}
}
