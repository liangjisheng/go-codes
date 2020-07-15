package main

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func createData() {
	db := connect()
	defer db.Close()

	u := &user{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	_, err := db.QueryOne(&user{}, "insert into users (name, emails) values (?name, ?emails) returning id", u)
	if err != nil {
		panic(err)
	}

	story1 := &story{
		Title:    "Cool story",
		AuthorID: u.ID,
	}
	_, err = db.QueryOne(&story{}, "insert into stories (title, author_id) values (?title, ?author_id) returning id", story1)
	if err != nil {
		panic(err)
	}

	u = &user{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	}
	_, err = db.QueryOne(&user{}, "insert into users (name, emails) values (?name, ?emails) returning id", u)
	if err != nil {
		panic(err)
	}
}

func query() {
	db := connect()
	defer db.Close()

	var users []user
	_, err := db.Query(&users, "select * from users")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)

	var u user
	_, err = db.QueryOne(&u, "select * from users where id = ?", 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	var users1 []user
	_, err = db.Query(&users1, "select * from users where id in (?)", pg.In([]int64{6, 7}))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users1)
}

// 通过Apply添加Where过滤条件查询
func apply() {
	db := connect()
	defer db.Close()

	// name := "admin"
	// name := "root"
	name := "asdf"

	filter := func(q *orm.Query) (*orm.Query, error) {
		if name == "admin" {
			q = q.Where("name = ?", "admin")
		}
		if name == "root" {
			q = q.Where("name = ?", "root")
		}
		return q, nil
	}

	var users []user
	err := db.Model(&users).Apply(filter).Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}

func query1() {
	db := connect()
	defer db.Close()

	var users []user

	db.Model(&users).Column("id", "name").
		OrderExpr("id ASC").Select()

	fmt.Println(users)
}

func query2() {
	db := connect()
	defer db.Close()

	var ids []int

	err := db.Model(&user{}).ColumnExpr("array_agg(id)").
		Select(pg.Array(&ids))
	if err != nil {
		panic(err)
	}

	fmt.Println(ids)
}

func whereGroup() {
	db := connect()
	defer db.Close()

	var users []user

	err := db.Model(&users).WhereGroup(func(query *orm.Query) (*orm.Query, error) {
		query = query.WhereOr("id=11").WhereOr("id=12")
		return query, nil
	}).Where("name is not null").Select()

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func whereIn() {
	db := connect()
	defer db.Close()

	var users []user

	err := db.Model(&users).
		WhereIn("id in (?)", []int{11, 12, 13}).Select()

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
