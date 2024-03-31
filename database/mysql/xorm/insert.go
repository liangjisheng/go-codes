package main

import (
	"fmt"
	"time"
	"xorm/xorm-models/models"
)

// Insert ...
func Insert() {
	// 新增数据
	doc3 := models.DoctorTb{
		Name:  "doctor-alice3",
		Age:   30,
		Ctime: time.Now(),
		Mtime: time.Now(),
	}
	i3, err := engine.InsertOne(doc3)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("insert result:", i3)
}
