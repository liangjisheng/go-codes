package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// 先执行create table 代码，成功后注释掉，再执行增删查改

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr("sql.Open", err)
	defer db.Close()

	// create table
	// stmt, err := db.Prepare(`CREATE TABLE userinfo (
	// 	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	username VARCHAR(64) NULL,
	// 	departname VARCHAR(64) NULL,
	// 	created DATE NULL);`)
	// checkErr("create userinfo", err)

	// res, err := stmt.Exec()
	// checkErr("create table userinfo", err)
	// fmt.Println("create table userinfo", res)

	// stmt, err = db.Prepare(`CREATE TABLE userdetail (
	// 	uid INT(10) NULL,
	// 	intro TEXT NULL,
	// 	profile TEXT NULL,
	// 	PRIMARY KEY (uid));`)
	// checkErr("create userdetail", err)

	// res, err = stmt.Exec()
	// checkErr("create table userdetail", err)
	// fmt.Println("create table userdetail", res)

	// insert
	stmt, err := db.Prepare("insert into userinfo(username, departname, created) values(?,?,?)")
	checkErr("insert", err)

	res, err := stmt.Exec("alice", "R&D", "2018-12-19")
	checkErr("insert", err)

	id, err := res.LastInsertId()
	checkErr("LastInsertId", err)
	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr("update", err)

	res, err = stmt.Exec("alice1", id)
	checkErr("update", err)

	affect, err := res.RowsAffected()
	checkErr("RowsAffected", err)
	fmt.Println(affect)

	// select
	rows, err := db.Query("select * from userinfo")
	checkErr("select", err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr("rows scan", err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr("delete", err)

	res, err = stmt.Exec(id)
	checkErr("delete", err)

	affect, err = res.RowsAffected()
	checkErr("RowsAffected", err)
	fmt.Println(affect)
}

func checkErr(str string, err error) {
	if err != nil {
		log.Println(str, err)
	}
}
