package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User define Model Struct
type User struct {
	ID   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:password@/alice?charset=utf8", 30)

	// 注册定义的model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// auto create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "alice"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "alice1"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{ID: user.ID}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete one
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
