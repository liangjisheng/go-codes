package mysql

import (
	"testing"
)

func TestUsers(t *testing.T) {
	db := Instance().db
	//可以使用 Table 方法临时指定表名
	// 根据 User 的字段创建 `deleted_users` 表
	db.Table("deleted_users").AutoMigrate(&User{})
}
