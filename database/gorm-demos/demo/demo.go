package mysql

import (
	"gorm.io/gorm"
	"time"
)

const (
	DemoTableName = "demo"
)

//在MySQL数据库中，也直到5.7这个版本，才开始引入JSON数据类型，在此之前如果想在表中保存JSON格式类型的数据，
//则需要依靠varchar或者text之类的数据类型，如果在低于5.7版本的数据库中使用了JSON类型来建表，显然是不会成功的
//JSON列存储的数据要么是NULL，要么必须是JSON格式数据，否则会报错
//JSON数据类型是没有默认值的（声明时"DEFAULT NULL"）

type Demo struct {
	ID          int64     `gorm:"column:id; primaryKey; AUTO_INCREMENT"`
	CreateTime  time.Time `gorm:"column:create_time; type:timestamp; default:CURRENT_TIMESTAMP; comment:'创建时间'"`
	CreateTime1 time.Time `gorm:"column:create_time_1; type:datetime; default:CURRENT_TIMESTAMP"`
	CreateTime2 int64     `gorm:"column:create_time_2; type:bigint; autoCreateTime"`
	JsonData    string    `gorm:"column:json_data; type:json; default:null;"`
	UpdateTime  time.Time `gorm:"column:update_time; type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	//autoUpdateTime:milli 使用时间戳毫秒数填充更新时间
	//autoUpdateTime:nano 使用时间戳填纳秒数充更新时间
	UpdateTime1 int64          `gorm:"column:update_time_1; type:bigint; autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp"`
}

func (Demo) TableName() string {
	return DemoTableName
}

func (s *Store) AddDemo(demo *Demo) error {
	var tmp Demo
	err := s.db.Table(DemoTableName).Where("`id` = ?", demo.ID).First(&tmp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err == gorm.ErrRecordNotFound {
		return s.db.Create(demo).Error
	} else {
		return s.db.Updates(demo).Error
	}
}

func (s *Store) GetDemo() ([]Demo, error) {
	res := make([]Demo, 0)
	err := s.db.Table(DemoTableName).Find(&res).Error
	if err != nil {
		return []Demo{}, err
	}
	return res, nil
}
