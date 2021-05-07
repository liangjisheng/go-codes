package mysql

import (
	"sync"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USER     = "root"
	PASSWORD = "ljs199711"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DATABASE = "study"
)

var (
	dbClient *DBClient
	once     sync.Once
)

// DBClient ..
type DBClient struct {
	db *gorm.DB
}

// NewDBClient ...
func NewDBClient() *DBClient {
	once.Do(func() {
		db, err := initDB()
		if err != nil {
			panic(err)
		}
		dbClient = &DBClient{
			db: db,
		}
	})

	return dbClient
}

func initDB() (*gorm.DB, error) {
	info := USER + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DATABASE + "?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	var err error
	db, err := gorm.Open("mysql", info)
	if err != nil {
		return nil, err
	}

	// 设置最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 设置最大打开连接数
	db.DB().SetMaxOpenConns(100)

	tables := []interface{}{
		&Admin{},
	}

	for _, table := range tables {
		db.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8mb4,COLLATE=utf8mb4_unicode_ci").AutoMigrate(table)
	}

	return db, nil
}
