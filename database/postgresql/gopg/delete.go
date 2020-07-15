package main

import "github.com/go-pg/pg"

func delete() {
	db := connect()
	defer db.Close()

	user1 := user{
		ID:     1,
		Name:   "user1",
		Emails: []string{"user1@qq.com", "user1@163.com"},
	}

	db.Delete(&user1)
}

func delete1() {
	db := connect()
	defer db.Close()

	user2 := user{
		ID:     2,
		Name:   "user2",
		Emails: []string{"user2@qq.com", "user2@163.com"},
	}
	user3 := user{
		ID:     3,
		Name:   "user3",
		Emails: []string{"user3@qq.com", "user3@163.com"},
	}

	db.Model(&user2, &user3).Delete()
}

func delete2() {
	db := connect()
	defer db.Close()

	ids := pg.In([]int{4, 5})
	db.Model(&user{}).Where("id IN (?)", ids).Delete()

	// 通过某个字段查找然后删除
	// db.Model(&User{}).Where("name=?","user2").Delete()
}
