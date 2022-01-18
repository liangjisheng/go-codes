package mysql

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	// AdminTableName ...
	AdminTableName = "admin"
)

// Admin ...
type Admin struct {
	ID       int64     `gorm:"column:id; primary_key; auto_increment"`
	User     string    `gorm:"column:user; not null; default:'ljs'; type:varchar(255); index:i_user; comment:'user name'"`
	Name     string    `gorm:"column:name; not null; unique:idx_uni_name; type:varchar(255)"`
	Password string    `gorm:"column:password; not null; default:18; type:char(255); index:i_password; comment:'user password'"`
	CreateAt int64     `gorm:"column:create_at; type:bigint; not null;"`
	T1       time.Time `gorm:"column:t1; autoCreateTime:nano"`
	T2       time.Time `gorm:"column:t2; autoUpdateTime:milli"`
	T3       time.Time `gorm:"column:t3; autoCreateTime"`
}

// TableName ...
func (Admin) TableName() string {
	return "admin"
}

// AddOrSaveAdmin ...
func (c *DBClient) AddOrSaveAdmin(admin *Admin) error {
	db := c.db.Table(AdminTableName)
	var tmp Admin
	err := db.Where("`user`= ?", admin.User).Find(&tmp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err == gorm.ErrRecordNotFound {
		return db.Create(admin).Error
	} else {
		admin.ID = tmp.ID
		return db.Save(admin).Error
	}
}

// GetUsers ...
func (c *DBClient) GetUsers() ([]Admin, error) {
	db := c.db.Table(AdminTableName)
	var res []Admin
	err := db.Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return res, nil
	}
}
