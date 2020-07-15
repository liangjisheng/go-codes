package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// GetConn ...
func GetConn() *sql.DB {
	// db, err := sql.Open("postgres", "user=root password=password dbname=root sslmode=disable")
	sqlURL := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	fmt.Println("----> get postgresql connection <----")
	db, err := sql.Open("postgres", sqlURL)
	checkErr(err, "-----> open datasources failed <-----")

	return db
}

func queryFunc(db *sql.DB) {
	selectByID := "select * from user"
	rows, err := db.Query(selectByID)
	checkErr(err, err.Error())

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		fmt.Printf("id: %d, name: %s, department: %s, created: %s\n",
			uid, username, department, created)
	}
}

func checkErr(e error, msg string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}
