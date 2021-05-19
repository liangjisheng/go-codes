package mysql

import "github.com/jinzhu/gorm"

const (
	// AdminTableName ...
	AdminTableName = "admin"
)

// Admin ...
type Admin struct {
	ID       int64  `json:"id" gorm:"primary_key; auto_increment; column:id"`
	User     string `json:"user" gorm:"not null; default:'ljs'; type:char(255); index:i_user; comment:'user name'"`
	Password string `json:"password" gorm:"not null; default:18; type:char(255); index:i_password; comment:'user password'"`
	CreateAt int64  `json:"create_at" gorm:"not null;"`
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
