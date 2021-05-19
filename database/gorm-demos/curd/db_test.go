package mysql

import (
	"fmt"
	"testing"
	"time"
)

func TestAddOrSaveAdmin(t *testing.T) {
	admin := &Admin{
		User:     "ljs",
		Password: "ljs",
		CreateAt: time.Now().Unix(),
	}

	err := NewDBClient().AddOrSaveAdmin(admin)
	if err != nil {
		fmt.Errorf("%+v", err)
		return
	}
	fmt.Println("AddOrSaveAdmin ok.")
}

func TestGetUsers(t *testing.T) {
	res, err := NewDBClient().GetUsers()
	if err != nil {
		fmt.Errorf("%+v", err)
		return
	}
	fmt.Printf("%+v\n", res)
}
