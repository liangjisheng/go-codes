package mysql

import "gorm.io/gorm"

const (
	ProductTableName = "product"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (Product) TableName() string {
	return ProductTableName
}
