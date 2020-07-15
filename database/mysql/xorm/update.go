package main

import (
	"fmt"
	"xorm/xorm-models/models"
)

// Update ...
func Update() {
	// 更新数据
	doc5 := models.DoctorTb{Name: "钟医生"}
	// 更新数据ID为7的记录名字更改为 "钟医生"
	iUpdate, _ := engine.Id(7).Update(&doc5)
	fmt.Printf("更新结果: %d\n\n", iUpdate)
}
