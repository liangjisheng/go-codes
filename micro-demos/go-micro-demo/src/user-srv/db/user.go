package db

import (
	"database/sql"
	"go-demos/micro-demos/go-micro-demo/src/user-srv/entity"
)

// SelectUserByID ...
func SelectUserByID(id int32) (*entity.User, error) {
	user := new(entity.User)
	if err := db.Get(user, "SELECT name, address, phone FROM user WHERE id = ?", id); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}
	return user, nil
}

// InsertUser ...
func InsertUser(user *entity.User) (int64, error) {
	rep, err := db.Exec("INSERT INTO `user`(`name`,`address`,`phone`) VALUE (?,?,?)", user.Name, user.Address, user.Phone)
	if err != nil {
		return 0, err
	}
	return rep.LastInsertId()
}

// ModifyUser ...
func ModifyUser(user *entity.User) error {
	_, err := db.Exec("UPDATE  `user` set `name` = ? ,`phone` = ?, `address` = ? where `id` = ?", user.Name, user.Phone, user.Address, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func DeleteUser(id int32) error {
	_, err := db.Exec("DELETE FROM `user` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	return nil
}
