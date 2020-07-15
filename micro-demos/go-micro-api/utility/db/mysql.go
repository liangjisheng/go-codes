package db

import (
	"time"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlDB *gorm.DB

// InitMysql ...
func InitMysql(dsn string, maxIdle, maxOpen int) (err error) {
	mysqlDB, err = gorm.Open("mysql", dsn)
	if err == nil {
		mysqlDB.DB().SetMaxIdleConns(maxIdle)
		mysqlDB.DB().SetMaxOpenConns(maxOpen)
		mysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	}
	return
}

// GetMysql 获取mysql连接
func GetMysql() *gorm.DB {
	return mysqlDB
}

// CloseMysql 关闭mysql
func CloseMysql() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}
