package main

import "fmt"

func insert() {
	db := connect()
	defer db.Close()

	// insert model
	user1 := &user{
		Name:   "user1",
		Emails: []string{"user1@qq.com", "user1@163.com"},
	}

	db.Insert(user1)
	fmt.Println("user1-->", user1)

	// insert models
	user2 := &user{
		Name:   "user2",
		Emails: []string{"user2@qq.com", "user2@163.com"},
	}

	user3 := &user{
		Name:   "user3",
		Emails: []string{"user3@qq.com", "user3@163.com"},
	}

	db.Insert(user2, user3)
	fmt.Println("user2-->", user2, "-user3-->", user3)

	// db.Model.insert
	user4 := &user{
		Name:   "user4",
		Emails: []string{"user4@qq.com", "user4@163.com"},
	}

	user5 := &user{
		Name:   "user5",
		Emails: []string{"user5@qq.com", "user5@163.com"},
	}

	db.Model(user4, user5).Insert()
	fmt.Println("user4-->", user4, "-user5-->", user5)

	// insert slice
	user6_7 := []user{
		{
			Name:   "user6",
			Emails: []string{"user6@qq.com", "user6@163.com"},
		}, {
			Name:   "user7",
			Emails: []string{"user7@qq.com", "user7@163.com"},
		},
	}

	db.Insert(&user6_7)
	fmt.Println("user6->", user6_7)
}
