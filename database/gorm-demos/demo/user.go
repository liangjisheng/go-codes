package mysql

import (
	"gorm.io/gorm"
	"time"
)

const (
	UserTableName = "user"
)

type User struct {
	ID       int64     `json:"id" gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserName string    `json:"username" gorm:"column:username; not null; type:varchar(128); default:''; uniqueIndex:idx_uni_name; comment:'user name'"`
	Password string    `json:"password" gorm:"column:password; not null; type:varchar(1024); default:''"`
	Email    string    `json:"email" gorm:"column:email; not null; type:varchar(128); default:''; index:idx_email"`
	UUID     string    `json:"uuid" gorm:"column:uuid; not null; type:varchar(64); default:''"`
	Deleted  int       `json:"deleted" gorm:"column:deleted; not null; type:tinyint(1); default:0"`
	CreateAt int64     `json:"create_at" gorm:"column:create_at; type:bigint; not null;"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at; type:datetime; not null; default:now()"`
}

func (u *User) TableName() string {
	return UserTableName
}

func (s *Store) AddOrSaveUser(user *User) error {
	var tmp User
	err := s.db.Model(&User{}).
		Where("`username` = ? or `email` = ?", user.UserName, user.Email).
		First(&tmp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err == gorm.ErrRecordNotFound {
		return s.db.Create(user).Error
	} else {
		user.ID = tmp.ID
		return s.db.Save(user).Error
	}
}

func (s *Store) GetUsers(start, limit int) ([]User, error) {
	res := make([]User, 0)
	err := s.db.Offset(start).Limit(limit).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return res, nil
	}
}

func (s *Store) GetUserByUsernameOrEmail(username, email string) (*User, error) {
	res := &User{}
	err := s.db.Model(&User{}).
		Where("`username` = ? or `email` = ?", username, email).First(res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return res, nil
	}
}
