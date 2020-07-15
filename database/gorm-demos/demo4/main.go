package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User ...
type User struct {
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

// TableName ...
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

var db *gorm.DB

func init() {
	username := "root"      //账号
	password := "ljs199711" //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "study"       //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 为了方便调试，了解gorm操作到底执行了怎么样的sql语句，开发的时候需要
	// 打开调试日志，这样gorm会打印出执行的每一条sql语句
	db.LogMode(true)

	//设置数据库连接池参数
	db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	// 使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
	// 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接
}

func main() {
	if db != nil {
		defer db.Close()
	}

	u := User{
		Username:   "tizi365",
		Password:   "123456",
		CreateTime: time.Now().Unix(),
	}

	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	var ids []int
	err := db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &ids).Error
	if err != nil {
		fmt.Println("select LAST_INSERT_ID() fail")
		return
	}
	fmt.Println("LAST_INSERT_ID", ids)
	// 提示：如果gorm设置了数据库连接池，那么每次执行数据库查询的时候都会从
	// 数据库连接池申请一个数据库连接，那么上述代码必须使用数据库事务，
	// 确保插入数据和查询自增id两条sql语句是在同一个数据库连接下执行，
	// 否则在高并发场景下，可能会查询不到自增id，或者查询到错误的id

	u = User{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	isNotFound := db.Where("username = ?", "tizi365").First(&u).RecordNotFound()
	if isNotFound {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Println(u.Username, u.Password)

	//更新
	//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'tizi365')
	db.Model(User{}).Where("username = ?", "tizi365").Update("password", "654321")

	//删除
	//自动生成Sql： DELETE FROM `users`  WHERE (username = 'tizi365')
	db.Where("username = ?", "tizi365").Delete(User{})
}
