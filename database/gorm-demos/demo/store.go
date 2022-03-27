package mysql

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	mysqlUserName = "root"
	mysqlPassword = "ljs199711"
	mysqlHost     = "127.0.0.1:3306"
	mysqlDatabase = "study"
)

var store *Store
var storeOnce sync.Once

type Store struct {
	db *gorm.DB
}

func Instance() *Store {
	storeOnce.Do(func() {
		gdb, err := initDb()
		if err != nil {
			panic(err)
		}
		store = &Store{
			db: gdb,
		}
	})
	return store
}

func initDb() (*gorm.DB, error) {
	//dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC",
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUserName, mysqlPassword, mysqlHost, mysqlDatabase)

	newLogger := gormLogger.Discard
	//newLogger = gormLogger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	gormLogger.Config{
	//		SlowThreshold: time.Second,     // 慢 SQL 阈值
	//		LogLevel:      gormLogger.Info, // Log level
	//		Colorful:      false,           // 禁用彩色打印
	//	},
	//)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	gdb, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	gdb.SetMaxIdleConns(100)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	gdb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	gdb.SetConnMaxLifetime(time.Hour)

	tables := []interface{}{
		&User{},
		&Demo{},
	}
	for _, table := range tables {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8mb4,COLLATE=utf8mb4_unicode_ci").AutoMigrate(table)
		if err != nil {
			log.Println(err)
		}
	}

	return db, nil
}
