package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type student struct {
	ID    int
	Name  string
	Age   int
	Sex   string
	Score float32
	Addr  string
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "test.db")
	orm.RegisterModel(new(student))
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	stu := new(student)
	stu.Name = "alice"
	stu.Age = 24
	stu.Sex = "m"
	stu.Score = 88
	stu.Addr = "shanxi.taiyuan"

	fmt.Println(o.Insert(stu))
}
