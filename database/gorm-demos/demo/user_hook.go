package mysql

import (
	"gorm.io/gorm"
	"log"
)

//BeforeCreate 定义钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate:", tx.Name())
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("AfterCreate:", tx.Name())
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	log.Println("AfterFind:", tx.Name())
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Println("BeforeUpdate:", tx.Name())
	return nil
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	log.Println("AfterUpdate:", tx.Name())
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println("BeforeDelete:", tx.Name())
	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	log.Println("AfterDelete:", tx.Name())
	return nil
}
