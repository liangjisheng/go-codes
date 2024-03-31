package models

import (
	"time"
)

type DoctorTb struct {
	Age   int       `xorm:"not null default 0 comment('age') INT"`
	Ctime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Id    int       `xorm:"not null pk autoincr INT"`
	Mtime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	Name  string    `xorm:"not null default '' comment('name') VARCHAR(50)"`
	Sex   int       `xorm:"not null default 0 comment('sex') INT"`
}
