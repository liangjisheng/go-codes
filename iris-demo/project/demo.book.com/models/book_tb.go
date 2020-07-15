package models

import (
	"time"
)

type BookTb struct {
	Addtime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Author       string    `xorm:"not null default '' comment('author') VARCHAR(100)"`
	Bookimage    string    `xorm:"not null default '' comment('图书封面') VARCHAR(300)"`
	Bookname     string    `xorm:"not null default '' comment('bookname') VARCHAR(100)"`
	Id           int       `xorm:"not null pk autoincr INT"`
	Introduction string    `xorm:"not null default '' comment('简介') VARCHAR(300)"`
	Press        string    `xorm:"not null default '' comment('出版社') VARCHAR(100)"`
	Price        string    `xorm:"not null default 0.00 comment('售价') DECIMAL(10,2)"`
	Publishtime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('出版时间') TIMESTAMP"`
	State        int       `xorm:"not null default 0 comment('状态') INT"`
	Updatetime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
