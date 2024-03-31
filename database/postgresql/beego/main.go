package main

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

// Student ...
type Student struct {
	ID   int64
	Name string
	Age  int
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=root password=password dbname=root host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterModel(new(Student))

	// auto create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	stu := new(Student)
	stu.Name = "alice"
	stu.Age = 25

	fmt.Println(o.Insert(stu))

	var stu1 Student
	err := o.QueryTable(&Student{}).Filter("age", 22).One(&stu1, "i_d", "name", "age")
	if err != nil {
		log.Println("===", err)
		return
	}
	fmt.Printf("%+v\n", stu1)
}
