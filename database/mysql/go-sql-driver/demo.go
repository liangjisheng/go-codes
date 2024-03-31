package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 使用database/sql接口对数据库表进行增删改查操作

func demo() {
	// user:password@/dbname
	db, err := sql.Open("mysql", "root:password@/alice?charset=utf8")
	checkErr(err)
	defer db.Close()

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("alice", "R&D", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("alice_update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// select
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}
