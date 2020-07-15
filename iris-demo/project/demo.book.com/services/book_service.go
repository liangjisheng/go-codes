package services

import (
	"demo.book.com/dao"
	"demo.book.com/dbsource"
	"demo.book.com/models"
)

// IBookService ...
type IBookService interface {
	GetList(query, sort string, pageSize int) []models.BookTb
	GetPageList(query, sort string, page, pageSize int) (int64, []models.BookTb)
	Get(id int) *models.BookTb
	Delete(id int) error
	Update(user *models.BookTb, columns []string) error
	Create(user *models.BookTb) error
}

type bookService struct {
	dao *dao.BookDao
}

// NewBookService ...
func NewBookService() IBookService {
	return &bookService{
		dao: dao.NewBookDao(dbsource.InstanceMaster()),
	}
}

func (b *bookService) GetList(query, sort string, pageSize int) []models.BookTb {
	return b.dao.GetList(query, sort, pageSize)
}

func (b *bookService) GetPageList(query, sort string, page, pageSize int) (int64, []models.BookTb) {
	return b.dao.GetPageList(query, sort, page, pageSize)
}

func (b *bookService) Get(id int) *models.BookTb {
	return b.dao.Get(id)
}

func (b *bookService) Delete(id int) error {
	_, err := b.dao.Delete(id)
	return err
}

func (b *bookService) Update(book *models.BookTb, columns []string) error {
	_, err := b.dao.Update(book, columns)
	return err
}

func (b *bookService) Create(book *models.BookTb) error {
	_, err := b.dao.Create(book)
	return err
}
