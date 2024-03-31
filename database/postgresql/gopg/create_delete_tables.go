package main

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func createTables(db *pg.DB) error {
	for _, model := range []interface{}{&user{}, &story{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteTable(db *pg.DB) error {
	err := db.DropTable(&user{}, &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
	return err
}

func createDeleteTables() {
	db := connect()
	defer db.Close()
	fmt.Println("db", db)

	err := createTables(db)
	if err != nil {
		panic(err)
	}

	// err = deleteTable(db)
	// if err != nil {
	// 	panic(err)
	// }
}
