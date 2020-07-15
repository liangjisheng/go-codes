package models

import (
	"go-demos/beego-demo/login-register/models/class"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(localhost:3306)/05blog?charset=utf8")
	orm.RegisterModel(new(class.User))
	orm.RunSyncdb("default", false, true)
}
