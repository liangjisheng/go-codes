package main

import "github.com/go-pg/pg"

const (
	host     = "localhost"
	port     = 5432
	username = "root"
	password = "password"
	dbname   = "root"
)

func connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: dbname,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "select 1")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	// createDeleteTables()
	// insert()
	// delete()
	// delete1()
	// delete2()
	// update()
	// update1()
	// createData()
	// query()
	// apply()
	// query1()
	// query2()
	// whereGroup()
	whereIn()
}
