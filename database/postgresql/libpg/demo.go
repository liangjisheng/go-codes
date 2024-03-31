package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

func demo() {
	db, err := sql.Open("postgres", "user=root password=password dbname=root sslmode=disable")
	checkErr(err)
	defer db.Close()

	// sqlCreateTables()
	sqlInsert(db)
	// sqlUpdate(db)
	// sqlQuery(db)
	// sqlDelete(db)
}

func sqlCreateTables() {
	db, err := sql.Open("postgres", "user=root password=password dbname=root sslmode=disable")
	checkErr(err)
	defer db.Close()

	// create table userinfo
	createTable, err := ioutil.ReadFile("create_userinfo.sql")
	checkErr(err)

	stmt, err := db.Prepare(string(createTable))
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(id)

	// create table userdetail
	createTable, err = ioutil.ReadFile("create_userdetail.sql")
	checkErr(err)

	stmt, err = db.Prepare(string(createTable))
	checkErr(err)

	res, err = stmt.Exec()
	checkErr(err)

	id, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(id)
}

func sqlInsert(db *sql.DB) {
	insert, err := ioutil.ReadFile("insert.sql")
	checkErr(err)

	stmt, err := db.Prepare(string(insert))
	checkErr(err)

	res, err := stmt.Exec("alice", "R&D", "2019-5-23")
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(id)
}

func sqlUpdate(db *sql.DB) {
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("aliceupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func sqlQuery(db *sql.DB) {
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func sqlDelete(db *sql.DB) {
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
