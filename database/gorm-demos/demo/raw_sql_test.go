package mysql

import (
	"database/sql"
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestRawSQL(t *testing.T) {
	db := Instance().db
	var user User

	type Result struct {
		ID   int
		Name string
		Age  int
	}

	//var result Result
	//db.Raw("SELECT id, name, age FROM users WHERE name = ?", "alice").Scan(&result)
	//t.Logf("%+v", result)

	//var age int
	//db.Raw("SELECT SUM(age) FROM users WHERE name = ?", "alice").Scan(&age)
	//t.Log("age ", age)

	//var users []User
	//err := db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "alice", 40).Scan(&users).Error
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Logf("%+v", users)

	db.Exec("DROP TABLE users")
	db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

	// Exec with SQL Expression
	db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "alice")

	//GORM 支持 sql.NamedArg、map[string]interface{}{} 或 struct 形式的命名参数

	db.Where("name1 = @name OR name2 = @name", sql.Named("name", "alice")).Find(&user)
	// SELECT * FROM `users` WHERE name1 = "alice" OR name2 = "alice"

	db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "alice"}).First(&user)
	// SELECT * FROM `users` WHERE name1 = "alice" OR name2 = "alice" ORDER BY `users`.`id` LIMIT 1

	// 原生 SQL 及命名参数
	db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
		sql.Named("name", "alice"), sql.Named("name2", "alice1")).Find(&user)
	// SELECT * FROM users WHERE name1 = "alice" OR name2 = "alice1" OR name3 = "alice"

	db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
		sql.Named("name", "alice"), sql.Named("name2", "alice1"))
	// UPDATE users SET name1 = "alice", name2 = "alice1", name3 = "alice"

	db.Raw("SELECT * FROM users WHERE (name1 = @name AND name3 = @name) AND name2 = @name2",
		map[string]interface{}{"name": "alice", "name2": "alice1"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "alice" AND name3 = "alice1") AND name2 = "alice1"

	type NamedArgument struct {
		Name  string
		Name2 string
	}

	db.Raw("SELECT * FROM users WHERE (name1 = @Name AND name3 = @Name) AND name2 = @Name2",
		NamedArgument{Name: "alice", Name2: "alice1"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "alice" AND name3 = "alice") AND name2 = "alice1"
}

func TestRawSQL1(t *testing.T) {
	db := Instance().db
	var user User

	//在不执行的情况下生成 SQL 及其参数，可以用于准备或测试生成的 SQL

	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = 1 ORDER BY `id`
	//stmt.Vars         //=> []interface{}{1}

	//ToSQL 返回生成的 SQL 但不执行。
	//GORM使用 database/sql 的参数占位符来构建 SQL 语句，它会自动转义参数以避免 SQL 注入，但我们不保证生成 SQL 的安全，请只用于调试
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]User{})
	})
	//sql //=> SELECT * FROM "users" WHERE id = 100 AND "users"."deleted_at" IS NULL ORDER BY age desc LIMIT 10
	t.Log(sql)

	//Row & Rows

	var name string
	var age int
	var email string
	// 使用 GORM API 构建 SQL
	row := db.Table("users").Where("name = ?", "alice").Select("name", "age").Row()
	row.Scan(&name, &age)

	// 使用原生 SQL
	row = db.Raw("select name, age, email from users where name = ?", "alice").Row()
	row.Scan(&name, &age, &email)

	// 使用 GORM API 构建 SQL
	rows, err := db.Model(&User{}).Where("name = ?", "alice").Select("name, age, email").Rows()
	defer rows.Close()
	t.Log(err)
	for rows.Next() {
		rows.Scan(&name, &age, &email)

		// 业务逻辑...
	}

	// 原生 SQL
	rows, err = db.Raw("select name, age, email from users where name = ?", "alice").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name, &age, &email)

		// 业务逻辑...
	}

	//使用 ScanRows 将一行记录扫描至 struct，例如
	rows, err = db.Model(&User{}).Where("name = ?", "alice").Select("name, age, email").Rows() // (*sql.Rows, error)
	defer rows.Close()

	for rows.Next() {
		// ScanRows 将一行扫描至 user
		db.ScanRows(rows, &user)

		// 业务逻辑...
	}

	//Connection Run multiple SQL in same db tcp connection (not in a transaction)
	db.Connection(func(tx *gorm.DB) error {
		tx.Exec("SET my.role = ?", "admin")

		tx.First(&User{})
		return nil
	})
}

//使用 Scopes 来动态指定查询的表

func TableOfYear(user *User, year int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := user.TableName() + strconv.Itoa(year)
		return db.Table(tableName)
	}
}

// Table form different database
func TableOfOrg(user *User, dbName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := dbName + "." + user.TableName()
		return db.Table(tableName)
	}
}

func TestScopes(t *testing.T) {
	db := Instance().db
	var user User
	var users []User

	db.Scopes(TableOfYear(&user, 2019)).Find(&users)
	// SELECT * FROM users_2019;

	db.Scopes(TableOfYear(&user, 2020)).Find(&users)
	// SELECT * FROM users_2020;

	db.Scopes(TableOfOrg(&user, "org1")).Find(&users)
	// SELECT * FROM org1.users;

	db.Scopes(TableOfOrg(&user, "org2")).Find(&users)
	// SELECT * FROM org2.users;
}
