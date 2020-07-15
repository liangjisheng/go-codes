package main

import "fmt"

func modify() {
	a1 := Admin{}
	// 查询单条记录 没有指定条件的时候,获取主键第一条记录
	db.Select([]string{"id", "user", "password"}).Where("id = ? AND user = ?", 1, "Test59").First(&a1)
	fmt.Println(a1)

	// 获取主键最后一条记录
	a2 := Admin{}
	db.Last(&a2)
	fmt.Println(a2)

	var admins []Admin
	db.Where("id > 20").Find(&admins)
	fmt.Println(admins)

	// 更新操作
	a3 := Admin{}
	a3.User = "Test-demo"
	a3.ID = 20
	a3.Password = a1.Password
	// save 更新或者保存
	db.Save(&a3)
	// 更新部分字段
	a4 := Admin{}
	a4.ID = 19
	//更新字段user
	db.Model(&a4).Update("user", "demo99")
	// 删除操作
	a5 := Admin{}
	a5.ID = 2
	errs := db.Delete(&a5).Error
	if errs == nil {
		fmt.Println("delete success")
	}
}
