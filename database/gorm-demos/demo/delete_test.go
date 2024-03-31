package mysql

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestDeleteUser(t *testing.T) {
	db := Instance().db

	//删除一条记录时，删除对象需要指定主键，否则会触发 批量 Delete
	user := User{
		ID: 3,
	}

	//DELETE from users where id = 3;
	err := db.Delete(&user).Error
	if err != nil {
		t.Error(err)
		return
	}

	// 带额外条件的删除
	db.Where("name = ?", "alice").Delete(&user)
	// DELETE from emails where id = 3 AND name = "alice";

	//GORM 允许通过主键(可以是复合主键)和内联条件来删除对象，它可以使用数字

	db.Delete(&User{}, 10)
	// DELETE FROM users WHERE id = 10;

	db.Delete(&User{}, "10")
	// DELETE FROM users WHERE id = 10;

	var users []User
	db.Delete(&users, []int{1, 2, 3})
	// DELETE FROM users WHERE id IN (1,2,3);

	//如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录

	db.Where("email LIKE ?", "%alice%").Delete(&User{})
	// DELETE from emails where email LIKE "%alice%";

	db.Delete(&User{}, "email LIKE ?", "%alice%")
	// DELETE from emails where email LIKE "%alice%";

	//如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	//对此，你必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式

	err = db.Delete(&User{}).Error // gorm.ErrMissingWhereClause

	db.Where("1 = 1").Delete(&User{})
	// DELETE FROM `users` WHERE 1=1

	db.Exec("DELETE FROM users")
	// DELETE FROM users

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
	// DELETE FROM users

	//返回被删除的数据，仅适用于支持 Returning 的数据库
	// 返回所有列
	db.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING *
	// users => []User{{ID: 1, Name: "alice", Role: "admin", Salary: 100}, {ID: 2, Name: "alice1", Role: "admin", Salary: 1000}}

	// 返回指定的列
	db.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
	// users => []User{{ID: 0, Name: "alice", Role: "", Salary: 100}, {ID: 0, Name: "alice1", Role: "", Salary: 1000}}
}

func TestSoftDeleteUser(t *testing.T) {
	db := Instance().db
	var user User

	//如果您的模型包含了一个 gorm.DeletedAt 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！
	//拥有软删除能力的模型调用 Delete 时，记录不会被数据库。但 GORM 会将 DeletedAt 置为当前时间， 并且你不能再通过普通的查询方法找到该记录

	// user 的 ID 是 `111`
	db.Delete(&user)
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("age = ?", 20).Delete(&User{})
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	// 在查询时会忽略被软删除的记录
	db.Where("age = 20").Find(&user)
	// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

	var users []User
	//可以使用 Unscoped 找到被软删除的记录
	db.Unscoped().Where("age = 20").Find(&users)
	// SELECT * FROM users WHERE age = 20;

	//也可以使用 Unscoped 永久删除匹配的记录
	db.Unscoped().Delete(&user)
	// DELETE FROM users WHERE id=10;
}
