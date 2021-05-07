package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product ...
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "./test.db")
	checkErr(err)
	defer db.Close()

	// Migrate the schema
	err = db.AutoMigrate(&Product{}).Error
	checkErr(err)

	// Create
	err = db.Create(&Product{Code: "L1212", Price: 1000}).Error
	checkErr(err)

	// Read
	var product Product
	// find product with id 1
	err = db.First(&product, 1).Error
	checkErr(err)

	// find product with code l1212
	err = db.First(&product, "code = ?", "L1212").Error
	checkErr(err)

	// Update - update product's price to 2000
	err = db.Model(&product).Update("Price", 2000).Error
	checkErr(err)

	// Delete - delete product
	err = db.Delete(&product).Error
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Panicln("db error")
	}
}
