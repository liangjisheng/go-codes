package dao

import (
	"demo.book.com/models"
	"github.com/go-xorm/xorm"
)

var tableName = "book_tb"

// BookDao ...
type BookDao struct {
	engine *xorm.Engine
}

// NewBookDao ...
func NewBookDao(engine *xorm.Engine) *BookDao {
	return &BookDao{
		engine: engine,
	}
}

// Get ...
func (b *BookDao) Get(id int) *models.BookTb {
	data := &models.BookTb{Id: id}
	ok, err := b.engine.Table(tableName).Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetList ...
func (b *BookDao) GetList(query, sort string, pageSize int) []models.BookTb {
	dataList := make([]models.BookTb, 0)

	session := b.engine.Table(tableName)
	if query != "" {
		session.Where(query)
	}
	if sort != "" {
		session.OrderBy(sort)
	}
	if pageSize > 0 {
		limit := pageSize
		start := 0
		session.Limit(limit, start)
	}
	err := session.Find(&dataList)
	if err != nil {
		return dataList
	}
	return dataList
}

// GetPageList ...
func (b *BookDao) GetPageList(query, sort string, page, pageSize int) (int64, []models.BookTb) {
	dataList := make([]models.BookTb, 0)

	session := b.engine.Table(tableName)
	if query != "" {
		session.Where(query)
	}
	if sort != "" {
		session.OrderBy(sort)
	}
	if pageSize > 0 {
		limit := pageSize
		start := page * pageSize
		session.Limit(limit, start)
	}
	total, err := session.FindAndCount(&dataList)
	if err != nil {
		return total, dataList
	}
	return total, dataList
}

// Delete ...
func (b *BookDao) Delete(id int) (int64, error) {
	data := &models.BookTb{Id: id}
	deleteID, err := b.engine.Table(tableName).Delete(&data)
	return deleteID, err
}

// Update ...
func (b *BookDao) Update(data *models.BookTb, columns []string) (int64, error) {
	updateID, err := b.engine.Table(tableName).ID(data.Id).MustCols(columns...).Update(data)
	return updateID, err
}

// Create ...
func (b *BookDao) Create(data *models.BookTb) (int64, error) {
	createID, err := b.engine.Table(tableName).Insert(data)
	return createID, err
}
