package mysql

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestUpdateSave(t *testing.T) {
	db := Instance().db
	var user User

	//Save 会保存所有的字段，即使字段是零值

	db.First(&user)

	user.Name = "alice1"
	user.Age = 100
	db.Save(&user)
	// UPDATE users SET name='alice1', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
}

func TestUpdateSingleColumn(t *testing.T) {
	db := Instance().db
	var user User

	//当使用 Update 更新单个列时，你需要指定条件，否则会返回 ErrMissingWhereClause 错误
	//使用了 Model 方法，且该对象主键有值，该值会被用于构建条件

	// 条件更新
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User 的 ID 是 `111`
	db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 根据条件和 model 的值进行更新
	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
}

func TestUpdateMultiColumn(t *testing.T) {
	db := Instance().db
	var user User

	//Updates 方法支持 struct 和 map[string]interface{} 参数。当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段

	// 根据 `struct` 更新属性，只会更新非零值的字段
	//db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 根据 `map` 更新属性
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	//注意 当通过 struct 更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
}

func TestUpdateSpecify(t *testing.T) {
	db := Instance().db
	var user User

	//如果您想要在更新时选定、忽略某些字段，您可以使用 Select、Omit

	// 使用 Map 进行 Select
	// User's ID is `111`:
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello' WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 使用 Struct 进行 Select（会 select 零值的字段）
	db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;

	// Select 所有字段（查询包括零值字段的所有字段）
	//db.Model(&user).Select("*").Update(User{Name: "alice", Role: "admin", Age: 0})

	// Select 除 Role 外的所有字段（包括零值字段的所有字段）
	//db.Model(&user).Select("*").Omit("Role").Update(User{Name: "alice", Role: "admin", Age: 0})
}

func TestBatchUpdate(t *testing.T) {
	db := Instance().db

	//如果您尚未通过 Model 指定记录的主键，则 GORM 会执行批量更新

	// 根据 struct 更新
	db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

	// 根据 map 更新
	db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

	//如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	//对此，你必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式，例如

	err := db.Model(&User{}).Update("name", "alice").Error // gorm.ErrMissingWhereClause
	if err != nil {
		t.Error(err)
		return
	}

	db.Model(&User{}).Where("1 = 1").Update("name", "alice")
	// UPDATE users SET `name` = "alice" WHERE 1=1

	db.Exec("UPDATE users SET name = ?", "alice")
	// UPDATE users SET name = "alice"

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "alice")
	// UPDATE users SET `name` = "alice"

	//获取受更新影响的行数

	// 通过 `RowsAffected` 得到更新的记录数
	result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

	t.Log(result.RowsAffected) // 更新的记录数
	t.Log(result.Error)        // 更新的错误
}

func TestUpdateExpr(t *testing.T) {
	db := Instance().db
	product := Product{
		Model: gorm.Model{ID: 3},
	}

	//GORM 允许使用 SQL 表达式更新列
	// product 的 ID 是 `3`
	db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

	db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

	db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;

	db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;
}

func TestUpdateSubSelect(t *testing.T) {
	db := Instance().db
	var user User
	type Company struct{}

	//使用子查询更新表
	db.Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
	// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

	db.Table("users as u").Where("name = ?", "alice").
		Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))

	db.Table("users as u").Where("name = ?", "alice").
		Updates(map[string]interface{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})
}

func TestUpdateColumn(t *testing.T) {
	db := Instance().db
	var user User

	//如果您想在更新时跳过 Hook 方法且不追踪更新时间，可以使用 UpdateColumn、UpdateColumns，其用法类似于 Update、Updates

	// 更新单个列
	db.Model(&user).UpdateColumn("name", "hello")
	// UPDATE users SET name='hello' WHERE id = 111;

	// 更新多个列
	db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE id = 111;

	// 更新选中的列
	db.Model(&user).Select("name", "age").UpdateColumns(User{Name: "hello", Age: 0})
	// UPDATE users SET name='hello', age=0 WHERE id = 111;
}

func TestUpdateReturn(t *testing.T) {
	db := Instance().db

	//返回被修改的数据，仅适用于支持 Returning 的数据库，例如

	// 返回所有列
	var users []User
	db.Model(&users).Clauses(clause.Returning{}).
		Where("name = ?", "alice").
		Update("age", gorm.Expr("age * ?", 2))
	// UPDATE `users` SET `salary`=salary * 2,`updated_at`="2021-10-28 17:37:23.19" WHERE role = "admin" RETURNING *
	// users => []User{{ID: 1, Name: "alice", Role: "admin", Salary: 100}, {ID: 2, Name: "alice1", Role: "admin", Salary: 1000}}
	t.Logf("%+v", users)

	// 返回指定的列
	db.Model(&users).Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "age"}}}).
		Where("name = ?", "alice").
		Update("age", gorm.Expr("age * ?", 2))
	// UPDATE `users` SET `salary`=salary * 2,`updated_at`="2021-10-28 17:37:23.19" WHERE role = "admin" RETURNING `name`, `salary`
	// users => []User{{ID: 0, Name: "alice", Role: "", Salary: 100}, {ID: 0, Name: "alice1", Role: "", Salary: 1000}}
}
