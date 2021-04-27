package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	user   = "root"
	pwd    = "ljs199711"
	dbip   = "127.0.0.1"
	dbport = "3306"
	dbname = "study"
)

// Admin ...
type Admin struct {
	ID       int64
	User     string
	Password string
}

// Finish ...
type Finish struct {
	ID         int    // GORM默认会使用ID字段作为表的主键
	Callid     string `gorm:"size:50;not null"` // 结构体标记,指定字段属性
	Subid      string `gorm:"size:100"`
	Aid        int64  `gorm:"index"`
	CreateTime time.Time
}

// Account ...
type Account struct {
	//gorm.Model 是一个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的结构体
	gorm.Model
	Appkey  string `gorm:"type:varchar(15);index:idx_appkey;not null"`
	Company string `gorm:"column:cpmpany_name;size:30"`
	Status  int8   `gorm:"default:1"` // 指定默认值
}

// TableName ...
func (Admin) TableName() string {
	return "vn_admin"
}

// TableName ...
func (Finish) TableName() string {
	return "vn_finish"
}

// TableName ...
func (Account) TableName() string {
	return "vn_account"
}

var db *gorm.DB

func init() {
	info := user + ":" + pwd + "@tcp(" + dbip + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	fmt.Println(info)
	var err error
	db, err = gorm.Open("mysql", info)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	} else {
		fmt.Println("mysql connect success")
	}

	// 设置最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 设置最大打开连接数
	db.DB().SetMaxOpenConns(100)

	if !db.HasTable("vn_admin") {
		// 运行给定模型的自动迁移，只会添加缺少的字段，不会删除/更改当前数据
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Admin{})
	}
	if !db.HasTable(&Finish{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Finish{})
	}
	if !db.HasTable(&Account{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Account{})
	}
}
