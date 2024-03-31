package mysql

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestGetUsers(t *testing.T) {
	res, err := Instance().GetUsers(1, 1)
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("res:", res)
}

func TestSelectSingleRecord(t *testing.T) {
	db := Instance().db

	var user User

	res := db.First(&user)
	//SELECT * FROM users ORDER BY id LIMIT 1;
	t.Log(res.Error)
	//检查 ErrRecordNotFound 错误
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {

	}
	t.Log(res.RowsAffected)
	t.Logf("%+v", user)

	// 获取一条记录，没有指定排序字段
	//db.Take(&user)
	//SELECT * FROM users LIMIT 1;
	//t.Logf("%+v", user)

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	//SELECT * FROM users ORDER BY id DESC LIMIT 1;
	t.Logf("%+v", user)
}

func TestSelectSingleRecord1(t *testing.T) {
	//如果主键是数字类型，您可以使用 内联条件 来检索对象。 传入字符串参数时，需要特别注意 SQL 注入问题，查看 安全 获取详情.
	db := Instance().db
	var user User
	var users []User

	db.First(&user, 1)
	// SELECT * FROM users WHERE id = 10;
	t.Logf("%+v", user)

	db.First(&user, "1")
	// SELECT * FROM users WHERE id = 10;
	t.Logf("%+v", user)

	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);
	t.Logf("%+v", users)

	//如果主键是字符串（例如像 uuid），查询将被写成这样：
	//db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	// 获取全部记录
	result := db.Find(&users)
	// SELECT * FROM users;

	t.Log(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	t.Log(result.Error)        // returns error
	t.Logf("%+v", users)
}

func TestSelectWhere(t *testing.T) {
	db := Instance().db
	var user User
	var users []User

	// 获取第一条匹配的记录
	db.Where("name = ?", "alice").First(&user)
	// SELECT * FROM users WHERE name = 'alice' ORDER BY id LIMIT 1;

	// 获取全部匹配的记录
	db.Where("name <> ?", "alice").Find(&users)
	// SELECT * FROM users WHERE name <> 'alice';

	// IN
	db.Where("name IN ?", []string{"alice", "alice1"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('alice','alice1');

	// LIKE
	db.Where("name LIKE ?", "%alice%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%alice%';

	// AND
	db.Where("name = ? AND age >= ?", "alice", 10).Find(&users)
	// SELECT * FROM users WHERE name = 'alice' AND age >= 10;

	lastWeek := "2022-04-10 00:00:00"
	today := "2022-04-16 00:00:00"
	// Time
	db.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2022-04-10 00:00:00';

	// BETWEEN
	db.Where("updated_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2022-04-10 00:00:00' AND '2022-04-16 00:00:00';
}

func TestSelectWhereStruct(t *testing.T) {
	db := Instance().db
	var user User
	var users []User

	// Struct
	db.Where(&User{Name: "alice", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "alice" AND age = 20 ORDER BY id LIMIT 1;
	t.Logf("%+v", user)

	// Map
	db.Where(map[string]interface{}{"name": "alice", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "alice" AND age = 20;
	t.Logf("%+v", user)

	// 主键切片条件
	db.Where([]int64{1, 2, 3}).Find(&users)
	// SELECT * FROM users WHERE id IN (1, 2, 3);
	t.Logf("%+v", users)

	//注意 当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件，例如：
	db.Where(&User{Name: "alice", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "alice";

	//如果想要包含零值查询条件，你可以使用 map，其会包含所有 key-value 的查询条件，例如
	db.Where(map[string]interface{}{"Name": "alice", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "alice" AND age = 0;

	//当使用 struct 进行查询时，你可以通过向 Where() 传入 struct 来指定查询条件的字段、值、表名，例如
	db.Where(&User{Name: "alice"}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE name = "alice" AND age = 0;

	db.Where(&User{Name: "alice"}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;
}

func TestSelectFind(t *testing.T) {
	//查询条件也可以被内联到 First 和 Find 之类的方法中，其用法类似于 Where
	db := Instance().db
	var user User
	var users []User

	// 根据主键获取记录，如果是非整型主键
	//db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// Plain SQL
	db.Find(&user, "name = ?", "alice")
	// SELECT * FROM users WHERE name = "alice";

	db.Find(&users, "name <> ? AND age > ?", "alice", 20)
	// SELECT * FROM users WHERE name <> "alice" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;
}

func TestSelectNot(t *testing.T) {
	db := Instance().db
	var user User
	var users []User

	db.Not("name = ?", "alice").First(&user)
	// SELECT * FROM users WHERE NOT name = "alice" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"alice", "alice1"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("alice", "alice1");

	// Struct
	db.Not(User{Name: "alice", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "alice" AND age <> 18 ORDER BY id LIMIT 1;

	// 不在主键切片中的记录
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
}

func TestSelectOr(t *testing.T) {
	db := Instance().db
	var users []User

	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'alice'").Or(User{Name: "alice1", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'alice' OR (name = 'alice1' AND age = 18);

	// Map
	db.Where("name = 'alice'").Or(map[string]interface{}{"name": "alice1", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'alice' OR (name = 'alice1' AND age = 18);
}

func TestSelectSpecifyField(t *testing.T) {
	db := Instance().db
	var users []User

	db.Select("name", "age").Find(&users)
	// SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age FROM users;

	//COALESCE(value,…)是一个可变参函数，可以使用多个参数
	//接受多个参数，返回第一个不为NULL的参数，如果所有参数都为NULL，此函数返回NULL；当它使用2个参数时，和IFNULL函数作用相同
	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	// SELECT COALESCE(age,'42') FROM users;
}

func TestSelectOrder(t *testing.T) {
	db := Instance().db
	var users []User

	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// 多个 order
	db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
}

func TestSelectLimitOffset(t *testing.T) {
	db := Instance().db
	var users []User
	var users1 []User
	var users2 []User

	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	// 通过 -1 消除 Limit 条件
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// 通过 -1 消除 Offset 条件
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)
}

type Result struct {
	name  string
	Total int
}

type ResultDate struct {
	Date  time.Time
	Total int
}

func TestSelectGroupHaving(t *testing.T) {
	db := Instance().db
	var result Result

	db.Model(&User{}).Select("name, sum(age) as total").
		Where("name LIKE ?", "group%").
		Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1

	db.Model(&User{}).Select("name, sum(age) as total").
		Group("name").Having("name = ?", "group").
		Find(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

	var resultDate ResultDate
	rows, err := db.Table("orders").
		Select("date(created_at) as date, sum(amount) as total").
		Group("date(created_at)").Rows()
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&resultDate)
	}

	rows, err = db.Table("orders").
		Select("date(created_at) as date, sum(amount) as total").
		Group("date(created_at)").
		Having("sum(amount) > ?", 100).Rows()
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&resultDate)
	}

	var results []ResultDate
	db.Table("orders").
		Select("date(created_at) as date, sum(amount) as total").
		Group("date(created_at)").
		Having("sum(amount) > ?", 100).Scan(&results)
}

func TestSelectDistinct(t *testing.T) {
	//db.Distinct("name", "age").Order("name, age desc").Find(&results)
}

func TestSelectJoin(t *testing.T) {
	db := Instance().db
	type result struct {
		Name  string
		Email string
	}

	var res result
	db.Model(&User{}).Select("users.name, emails.email").
		Joins("left join emails on emails.user_id = users.id").
		Scan(&res)
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

	rows, err := db.Table("users").
		Select("users.name, emails.email").
		Joins("left join emails on emails.user_id = users.id").
		Rows()
	if err != nil {
		t.Error(err)
		return
	}

	defer rows.Close()

	for rows.Next() {

	}

	var results []result
	db.Table("users").Select("users.name, emails.email").
		Joins("left join emails on emails.user_id = users.id").
		Scan(&results)

	var user User
	// 带参数的多表连接
	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "alice@example.org").
		Joins("JOIN credit_cards ON credit_cards.user_id = users.id").
		Where("credit_cards.number = ?", "411111111111").Find(&user)
}

func TestSelectScan(t *testing.T) {
	db := Instance().db
	type Result struct {
		Name string
		Age  int
	}

	var result Result
	db.Table("users").Select("name", "age").
		Where("name = ?", "Antonio").Scan(&result)

	// Raw SQL
	db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
}

func TestSelectLocking(t *testing.T) {
	db := Instance().db
	var users []User

	db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
	// SELECT * FROM `users` FOR UPDATE

	db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Find(&users)
	// SELECT * FROM `users` FOR SHARE OF `users`

	db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	}).Find(&users)
	// SELECT * FROM `users` FOR UPDATE NOWAIT
}

func TestAdvancedSelect(t *testing.T) {
	db := Instance().db
	var user User
	var users []User
	var orders []User
	var results []User

	//子查询可以嵌套在查询中，GORM 允许在使用 *gorm.DB 对象作为参数时生成子查询
	db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
	// SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");

	subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
	db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
	// SELECT AVG(age) as avgage FROM `users` GROUP BY `name` HAVING AVG(age) > (SELECT AVG(age) FROM `users` WHERE name LIKE "name%")

	//From 子查询
	db.Table("(?) as u", db.Model(&User{}).Select("name", "age")).Where("age = ?", 18).Find(&User{})
	// SELECT * FROM (SELECT `name`,`age` FROM `users`) as u WHERE `age` = 18

	subQuery1 := db.Model(&User{}).Select("name")
	subQuery2 := db.Model(&Demo{}).Select("name")
	db.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&User{})
	// SELECT * FROM (SELECT `name` FROM `users`) as u, (SELECT `name` FROM `demos`) as p

	//使用 Group 条件可以更轻松的编写复杂 SQL
	_ = db.Where(
		db.Where("pizza = ?", "pepperoni").Where(db.Where("size = ?", "small").Or("size = ?", "medium")),
	).Or(
		db.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
	).Find(&User{}).Statement
	// SELECT * FROM `pizzas` WHERE (pizza = "pepperoni" AND (size = "small" OR size = "medium")) OR (pizza = "hawaiian" AND size = "xlarge")

	//带多个列的 in 查询
	db.Where("(name, age, role) IN ?", [][]interface{}{{"alice", 18, "admin"}, {"alice1", 19, "user"}}).Find(&users)
	// SELECT * FROM users WHERE (name, age, role) IN (("alice", 18, "admin"), ("alice1", 19, "user"));

	//GORM 支持 sql.NamedArg 和 map[string]interface{}{} 形式的命名参数，例如
	db.Where("name1 = @name OR name2 = @name", sql.Named("name", "alice")).Find(&user)
	// SELECT * FROM `users` WHERE name1 = "alice" OR name2 = "alice"

	db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "alice"}).First(&user)
	// SELECT * FROM `users` WHERE name1 = "alice" OR name2 = "alice" ORDER BY `users`.`id` LIMIT 1

	//GORM 允许扫描结果至 map[string]interface{} 或 []map[string]interface{}，此时别忘了指定 Model 或 Table，例如
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result, "id = ?", 1)

	var results1 []map[string]interface{}
	db.Table("users").Find(&results1)
}

func TestAdvancedFirstOrInit(t *testing.T) {
	db := Instance().db
	var user User

	//获取第一条匹配的记录，或者根据给定的条件初始化一个实例（仅支持 sturct 和 map 条件）

	// 未找到 user，则根据给定的条件初始化一条记录
	db.FirstOrInit(&user, User{Name: "non_existing"})
	// user -> User{Name: "non_existing"}

	// 找到了 `name` = `alice` 的 user
	db.Where(User{Name: "alice"}).FirstOrInit(&user)
	// user -> User{ID: 111, Name: "alice", Age: 18}

	// 找到了 `name` = `alice` 的 user
	db.FirstOrInit(&user, map[string]interface{}{"name": "alice"})
	// user -> User{ID: 111, Name: "alice", Age: 18}

	//如果没有找到记录，可以使用包含更多的属性的结构体初始化 user，Attrs 不会被用于生成查询 SQL

	// 未找到 user，则根据给定的条件以及 Attrs 初始化 user
	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
	// user -> User{Name: "non_existing", Age: 20}

	// 未找到 user，则根据给定的条件以及 Attrs 初始化 user
	db.Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = 'non_existing' ORDER BY id LIMIT 1;
	// user -> User{Name: "non_existing", Age: 20}

	// 找到了 `name` = `alice` 的 user，则忽略 Attrs
	db.Where(User{Name: "alice"}).Attrs(User{Age: 20}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = alice' ORDER BY id LIMIT 1;
	// user -> User{ID: 111, Name: "alice", Age: 18}

	//不管是否找到记录，Assign 都会将属性赋值给 struct，但这些属性不会被用于生成查询 SQL，也不会被保存到数据库

	// 未找到 user，根据条件和 Assign 属性初始化 struct
	db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
	// user -> User{Name: "non_existing", Age: 20}

	// 找到 `name` = `alice` 的记录，依然会更新 Assign 相关的属性
	db.Where(User{Name: "alice"}).Assign(User{Age: 200}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = alice' ORDER BY id LIMIT 1;
	// user -> User{ID: 111, Name: "alice", Age: 200}
	t.Logf("%+v", user)
}

func TestAdvancedFirstOrCreate(t *testing.T) {
	db := Instance().db
	var user User

	//获取第一条匹配的记录，或者根据给定的条件创建一条新纪录（仅支持 sturct 和 map 条件）

	// 未找到 user，则根据给定条件创建一条新纪录
	db.FirstOrCreate(&user, User{Name: "non_existing"})
	// INSERT INTO "users" (name) VALUES ("non_existing");
	// user -> User{ID: 112, Name: "non_existing"}

	// 找到了 `name` = `alice` 的 user
	db.Where(User{Name: "alice"}).FirstOrCreate(&user)
	// user -> User{ID: 111, Name: "alice", "Age": 18}

	//如果没有找到记录，可以使用包含更多的属性的结构体创建记录，Attrs 不会被用于生成查询 SQL

	// 未找到 user，根据条件和 Assign 属性创建记录
	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
	// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	// user -> User{ID: 112, Name: "non_existing", Age: 20}

	// 找到了 `name` = `alice` 的 user，则忽略 Attrs
	db.Where(User{Name: "alice"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'alice' ORDER BY id LIMIT 1;
	// user -> User{ID: 111, Name: "alice", Age: 18}

	//不管是否找到记录，Assign 都会将属性赋值给 struct，并将结果写回数据库

	// 未找到 user，根据条件和 Assign 属性创建记录
	db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1;
	// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	// user -> User{ID: 112, Name: "non_existing", Age: 20}

	// 找到了 `name` = `alice` 的 user，依然会根据 Assign 更新记录
	db.Where(User{Name: "alice"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'alice' ORDER BY id LIMIT 1;
	// UPDATE users SET age=20 WHERE id = 111;
	// user -> User{ID: 111, Name: "alice", Age: 20}
}

func TestSelectFindInBatches(t *testing.T) {
	db := Instance().db
	var results []User

	// 每次批量处理 100 条
	result := db.Where("processed = ?", false).FindInBatches(&results, 100, func(tx *gorm.DB, batch int) error {
		for _, result := range results {
			// 批量处理找到的记录
			t.Log(result)
		}

		tx.Save(&results)

		t.Log(tx.RowsAffected) // 本次批量操作影响的记录数

		//batch 1, 2, 3

		// 如果返回错误会终止后续批量操作
		return nil
	})

	t.Log(result.Error)        // returned error
	t.Log(result.RowsAffected) // 整个批量操作影响的记录数
}

func TestSelectPluck(t *testing.T) {
	db := Instance().db
	var users []User

	//Pluck 用于从数据库查询单个列，并将结果扫描到切片。如果您想要查询多列，您应该使用 Select 和 Scan

	var ages []int64
	db.Model(&users).Pluck("age", &ages)

	var names []string
	db.Model(&User{}).Pluck("name", &names)

	db.Table("deleted_users").Pluck("name", &names)

	// Distinct Pluck
	db.Model(&User{}).Distinct().Pluck("Name", &names)
	// SELECT DISTINCT `name` FROM `users`

	// 超过一列的查询，应该使用 `Scan` 或者 `Find`，例如：
	db.Select("name", "age").Scan(&users)
	db.Select("name", "age").Find(&users)
}

func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}

func OrderStatus(status []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status IN (?)", status)
	}
}

func TestSelectScope(t *testing.T) {
	//Scopes 允许你指定常用的查询，可以在调用方法时引用这些查询
	db := Instance().db
	var orders []User

	db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
	// 查找所有金额大于 1000 的信用卡订单

	db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
	// 查找所有金额大于 1000 的货到付款订单

	db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
	// 查找所有金额大于 1000 且已付款或已发货的订单
}

func TestSelectCount(t *testing.T) {
	db := Instance().db

	var count int64
	db.Model(&User{}).Where("name = ?", "alice").Or("name = ?", "alice1").Count(&count)
	// SELECT count(1) FROM users WHERE name = 'alice' OR name = 'alice1'

	db.Model(&User{}).Where("name = ?", "alice").Count(&count)
	// SELECT count(1) FROM users WHERE name = 'alice'; (count)

	db.Table("deleted_users").Count(&count)
	// SELECT count(1) FROM deleted_users;

	// Count with Distinct
	db.Model(&User{}).Distinct("name").Count(&count)
	// SELECT COUNT(DISTINCT(`name`)) FROM `users`

	db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	// SELECT count(distinct(name)) FROM deleted_users

	// Count with Group
	//users := []User{
	//	{Name: "name1"},
	//	{Name: "name2"},
	//	{Name: "name3"},
	//	{Name: "name3"},
	//}

	db.Model(&User{}).Group("name").Count(&count)
	//count // => 3
}
