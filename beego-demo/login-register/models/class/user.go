package class

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// User ...
type User struct {
	ID       int `orm:"pk"`
	Password string
}

// ReadDB select a user
func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return err
}

// Create insert a user
func (u *User) Create() (err error) {
	o := orm.NewOrm()
	fmt.Println("Create user success!")
	_, _ = o.Insert(u)
	return err
}

// Update update user
func (u *User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	return err
}
