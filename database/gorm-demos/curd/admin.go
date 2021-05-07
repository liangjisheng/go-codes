package mysql

import "github.com/jinzhu/gorm"

const (
	AdminTableName = "admin"
)

// Admin ...
type Admin struct {
	ID       int64
	User     string	`gorm:"default:'ljs'"`
	Password string `gorm:"default:18"`
	CreateAt int64
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
