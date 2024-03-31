package mysql

import (
	"testing"
	"time"
)

func TestAddDemo(t *testing.T) {
	demo := &Demo{
		CreateTime:  time.Now(),
		CreateTime1: time.Now(),
		UpdateTime:  time.Now(),
	}

	err := Instance().AddDemo(demo)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("AddDemo ok.")
}

func TestDeleteDemo(t *testing.T) {
	v := &Demo{
		ID: 2,
	}
	Instance().db.Model(&Demo{}).Delete(v)

	//Instance().db.Model(&Demo{}).Delete("id = ?", 12)
}

func TestGetDemo(t *testing.T) {
	res, err := Instance().GetDemo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", res)
}

func TestProduct(t *testing.T) {
	//Instance().db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	Instance().db.First(&product, 1) // 根据整型主键查找
	t.Logf("%+v", product)
	//Instance().db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//t.Logf("%+v", product)

	// Update - 将 product 的 price 更新为 200
	//Instance().db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//Instance().db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//Instance().db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	Instance().db.Delete(&product, 1)
}

func TestSoftDelete(t *testing.T) {
	user := User{Name: "alice", Age: 20}
	db := Instance().db

	//var count int64
	//var age int

	//if db.Model(&User{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 1 {
	//	t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
	//}

	//if db.Model(&User{}).Select("age").Where("name = ?", user.Name).Scan(&age).Error != nil || age != user.Age {
	//	t.Errorf("Age soft deleted record, expects: %v, got: %v", 0, age)
	//}

	if err := db.Delete(&user, "name = ?", user.Name).Error; err != nil {
		t.Fatalf("No error should happen when soft delete user, but got %v", err)
	}

	//if user.DeletedAt == 0 {
	//	t.Errorf("user's deleted at should not be zero, DeletedAt: %v", user.DeletedAt)
	//}

	//sql := db.Session(&gorm.Session{DryRun: true}).
	//	Delete(&user, "`user`.`id` = ?", 1).Statement.SQL.String()
	//if !regexp.MustCompile(`UPDATE .user. SET .deleted_at.=.* WHERE .user.\..id. = .* AND .user.\..deleted_at. = ?`).MatchString(sql) {
	//	t.Fatalf("invalid sql generated, got %v", sql)
	//}

	//if db.First(&User{}, "name = ?", user.Name).Error == nil {
	//	t.Errorf("Can't find a soft deleted record")
	//}

	//count = 0
	//if db.Model(&User{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 0 {
	//	t.Errorf("Count soft deleted record, expects: %v, got: %v", 0, count)
	//}

	//age = 0
	//if err := db.Model(&User{}).Select("age").Where("name = ?", user.Name).Scan(&age).Error; err != nil || age != 0 {
	//	t.Fatalf("Age soft deleted record, expects: %v, got: %v, err %v", 0, age, err)
	//}

	//if err := db.Unscoped().First(&User{}, "name = ?", user.Name).Error; err != nil {
	//	t.Errorf("Should find soft deleted record with Unscoped, but got err %s", err)
	//}

	//count = 0
	//if db.Unscoped().Model(&User{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 1 {
	//	t.Errorf("Count soft deleted record, expects: %v, count: %v", 1, count)
	//}

	//age = 0
	//if db.Unscoped().Model(&User{}).Select("age").Where("name = ?", user.Name).Scan(&age).Error != nil || age != user.Age {
	//	t.Errorf("Age soft deleted record, expects: %v, got: %v", 0, age)
	//}

	//db.Unscoped().Delete(&user, "name = ?", user.Name)
	//if err := db.Unscoped().First(&User{}, "name = ?", user.Name).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	t.Errorf("Can't find permanently deleted record")
	//}
}
