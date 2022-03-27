package mysql

import (
	"fmt"
	"testing"
	"time"
)

func TestAddOrSaveUser(t *testing.T) {
	for i := 0; i < 2; i++ {
		user := &User{
			UserName: fmt.Sprintf("%d", i),
			Password: fmt.Sprintf("%d", i),
			Email:    fmt.Sprintf("%d", i),
			UUID:     GetUUID(),
			Deleted:  0,
			CreateAt: time.Now().Unix(),
			UpdateAt: time.Now(),
		}

		err := Instance().AddOrSaveUser(user)
		if err != nil {
			t.Error("err:", err)
			continue
		}
		t.Log("AddOrSaveUser ok.")
	}
}

func TestGetUsers(t *testing.T) {
	res, err := Instance().GetUsers(1, 1)
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("res:", res)
}

func TestGetUserByUsername(t *testing.T) {
	res, err := Instance().GetUserByUsernameOrEmail("ljs", "")
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("res:", res)
}
