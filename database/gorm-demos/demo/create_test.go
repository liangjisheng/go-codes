package mysql

import (
	"testing"
	"time"

	"gorm.io/gorm/clause"
)

func TestAddOrSaveUser(t *testing.T) {
	age := 0
	user := User{
		Name: "alice",
		Age1: &age,
	}

	err := Instance().AddOrSaveUser(&user)
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("AddOrSaveUser ok.")
}

func TestCreateUserClauses(t *testing.T) {
	//GORM 为不同数据库提供了兼容的 Upsert 支持
	user := User{
		Name: "alice",
		Age:  20,
	}

	// 在冲突时，什么都不做
	Instance().db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// 在冲突时，更新除主键以外的所有列到新值。
	//Instance().db.Clauses(clause.OnConflict{
	//	UpdateAll: true,
	//}).Create(&users)
}

func TestBatchInsert(t *testing.T) {
	users := []User{
		{
			Name: "1",
			Age:  1,
		},
		{
			Name: "2",
			Age:  2,
		},
	}

	//2种批量 insert 的方式
	//err := Instance().db.Create(&users).Error
	err := Instance().db.CreateInBatches(&users, len(users)).Error

	if err != nil {
		t.Error(err)
		return
	}

	for _, user := range users {
		t.Log(user.ID)
	}
}

func TestCreateUserMap(t *testing.T) {
	//根据 map 创建记录时，association,hook 不会被调用，且主键也不会自动填充

	Instance().db.Model(&User{}).Create(map[string]interface{}{
		"name": "alice", "age": 18, "created_at": time.Now().Unix(),
	})

	// batch insert from `[]map[string]interface{}{}`
	Instance().db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "alice1", "Age": 18, "created_at": time.Now().Unix()},
		{"Name": "alice2", "Age": 20, "created_at": time.Now().Unix()},
	})
}

func TestCreateUserSpecifyField(t *testing.T) {
	user := User{
		Name: "alice",
		Age:  20,
	}
	//这样会导致其他字段是默认值, 比如 deleted_at
	Instance().db.Select("name", "age").Create(&user)
	//INSERT INTO `users` (`name`,`age`) VALUES ("alice", 20);
}
