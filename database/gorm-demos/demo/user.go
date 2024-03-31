package mysql

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

const (
	UserTableName = "users"
)

type User struct {
	ID int64 `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	//uniqueIndex:idx_uni_name;
	Name string `gorm:"column:name; not null; type:varchar(128); default:''; comment:'user name'"`
	//使用值类型, create,update 时如果 go struct 的字段是默认值, 则 gorm 不会使用 go 的默认值, 而是会使用 default 定义的默认值
	Age int `gorm:"column:age; not null; type:smallint; check:age >= 0; default:1"`
	//使用指针可以解决上面的问题
	Age1 *int `gorm:"column:age1; not null; type:smallint; check:age >= 0; default:1"`
	//使用 sql.NullInt32 也可以解决 go struct 字段默认值的 create,update 问题
	Age2      sql.NullInt32 `gorm:"column:age2; type:smallint; default:0"`
	CreatedAt int64         `gorm:"column:created_at; type:bigint; not null; autoCreateTime"`
	UpdatedAt time.Time     `gorm:"column:updated_at; type:datetime; not null; default:now() ON UPDATE CURRENT_TIMESTAMP"`
	//yyyy-mm-dd hh:mm:ss
	//DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; type:timestamp"`
	//gorm:"softDelete:milli" or gorm:"softDelete:nano" 支持毫秒或者纳秒
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at; type:bigint; not null; default:0"`
}

//TableName 不支持动态变化，它会被缓存下来以便后续使用。想要使用动态表名，你可以使用 Scopes

func (User) TableName() string {
	return UserTableName
}

func UserTable(user User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		//if user.Admin {
		//	return tx.Table("admin_users")
		//}

		//return tx.Table("users")

		return nil
	}
}

//db.Scopes(UserTable(user)).Create(&user)

func (s *Store) AddOrSaveUser(user *User) error {
	var tmp User
	err := s.db.Model(&User{}).Where("name = ?", user.Name).First(&tmp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err == gorm.ErrRecordNotFound {
		//user.ID             //返回插入数据的主键
		//result.Error        // 返回 error
		//result.RowsAffected // 返回插入记录的条数
		return s.db.Model(&User{}).Create(user).Error
	} else {
		return s.db.Model(&User{}).Where("name = ?", user.Name).Updates(user).Error
	}
}

func (s *Store) GetUsers(start, limit int) ([]User, error) {
	res := make([]User, 0)
	err := s.db.Model(&User{}).Offset(start).Limit(limit).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return res, nil
	}
}

func (s *Store) UserAgeAddOne(name string) error {
	return s.db.Model(&User{}).
		Where("name = ?", name).
		Update("age", gorm.Expr("age + ?", 1)).Error
}
