package main

func update() {
	db := connect()
	defer db.Close()

	u := user{
		ID: 6,
	}

	err := db.Select(&u)
	if err != nil {
		panic(err)
	}

	u.Name = "user6_"
	err = db.Update(&u)
	if err != nil {
		panic(err)
	}
}

func update1() {
	db := connect()
	defer db.Close()

	u := &user{
		ID:     6,
		Name:   "user66",
		Emails: []string{"user6@qq.com", "user6@163.com"},
	}

	db.Model(u).Set("name=?name,emails=?emails").Where("id=?id").Update()

	// 利用主键update
	// db.Model(user).Column("name","emails").WherePK().Update()
}
