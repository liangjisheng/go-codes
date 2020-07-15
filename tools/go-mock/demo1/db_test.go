package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	// m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	// m.EXPECT().Get(gomock.Any()).Do(func(key string) {
	// 	t.Log(key)
	// })
	// m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
	// 	if key == "Sam" {
	// 		return 630, nil
	// 	}
	// 	return 0, errors.New("not exist")
	// })

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}

func TestGetFromDB1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	if v := GetFromDB(m, "ABC"); v != -1 {
		t.Fatal("expect -1, but got", v)
	}
	if v := GetFromDB(m, "DEF"); v != -1 {
		t.Fatal("expect -1, but got", v)
	}
}

func TestGetFromDB2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)
	if v := GetFromDB(m, "Tom"); v != 0 {
		t.Fatal("expect 0, but got", v)
	}
	if v := GetFromDB(m, "Sam"); v != 630 {
		t.Fatal("expect 630, but got", v)
	}
}
