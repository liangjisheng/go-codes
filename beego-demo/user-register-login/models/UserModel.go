package models

import (
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// User ...
type User struct {
	ID       int    `form:"-" orm:"column(id)" json:"id"`
	UserName string `form:"username" orm:"column(username)" json:"username"`
	Age      int    `form:"age" json:"age"`
	Sex      string `form:"sex" json:"sex"`
	Mobile   string `form:"mobile" json:"mobile"`
	Password string `json:"password"`
	Email    string `form:"email" json:"email"`
}

// Result ...
type Result struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}

// QueryUserByID ...
func QueryUserByID(id int) *User {
	var user User
	orm := orm.NewOrm()
	orm.QueryTable("user").Filter("id", id).One(&user, "username", "age", "sex", "mobile")
	logs.Info(">>>> query user by user id from database <<<<")
	return &user
}

// QueryUserList ...
func QueryUserList() []*User {
	var users []*User
	orm := orm.NewOrm()
	orm.QueryTable("user").All(&users, "username", "age", "sex", "mobile")
	return users
}

// InsertUser ...
func InsertUser(u *User) int64 {
	orm := orm.NewOrm()
	id, err := orm.Insert(u)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

// QueryByNamePwd ...
func QueryByNamePwd(username, password string) bool {
	logs.Info(">>>> query user by user name and password from database <<<<")

	var user User
	o := orm.NewOrm()
	err := o.QueryTable(&User{}).Filter("username", username).Filter("password", password).One(&user, "username", "age", "sex", "mobile", "email")

	var result bool
	if err != nil {
		logs.Error(err)
		result = false
	} else {
		result = true
	}
	return result
}
