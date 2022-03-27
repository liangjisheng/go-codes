package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB

// Setup ...
func Setup() {
	var err error
	DB, err = gorm.Open("mysql", "user:pass@/study?charset=utf8&parseTime=True&loc=Local&timeout=100ms")
	if err != nil {
		fmt.Printf("mysql connect err: %+v", err)
		os.Exit(-1)
	}
	if DB.Error != nil {
		fmt.Printf("database err: %+v", DB.Error)
		os.Exit(-1)
	}

	AutoMigrateAll()
}

// AutoMigrateAll ...
func AutoMigrateAll() {
	DB.AutoMigrate(&User{})
}
