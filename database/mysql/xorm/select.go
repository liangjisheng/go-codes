package main

import (
	"fmt"
	"xorm/xorm-models/models"
)

// Select ...
func Select() {
	// 查询单条数据
	var doc models.DoctorTb
	b, _ := engine.Where("name=?", "doctor-alice").Get(&doc)
	if b {
		fmt.Printf("doc: %+v\n", doc)
	} else {
		fmt.Println("data not exist")
	}

	// 查询单条数据方式2 会根据结构体的
	doc2 := models.DoctorTb{Name: "doctor-alice1"}
	b, _ = engine.Get(&doc2)
	fmt.Printf("doc2: %+v\n", doc2)

	// 查询列表
	docList := make([]models.DoctorTb, 0)
	engine.Where("age > ? or name like ?", 25, "doctor%").Find(&docList)
	fmt.Printf("docList：%+v\n\n", docList)

	// 查询列表方式2
	docList2 := make([]models.DoctorTb, 0)
	engine.Where("age > ?", 25).Or("name like ?", "doctor%").OrderBy("Id desc").Find(&docList2)
	fmt.Printf("docList2: %+v\n\n", docList2)

	// 查询分页
	docList3 := make([]models.DoctorTb, 0)
	page := 0     // 页索引
	pageSize := 2 // 每页数据
	limit := pageSize
	start := page * pageSize
	totalCount, err := engine.Where("age > ? or name like ?", 25, "doctor%").Limit(limit, start).FindAndCount(&docList3)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("总记录数：", totalCount, "docList3：", docList3)
	fmt.Println()

	// 直接用语句查询
	docList4 := make([]models.DoctorTb, 0)
	engine.SQL("select * from doctor_tb where age > ?", 25).Find(&docList4)
	fmt.Printf("docList4: %+v\n\n", docList4)

	// 指定表名查询.Table()
	user := models.UserTb{Id: 67}
	b, _ = engine.Table("user_tb").Get(&user)
	fmt.Printf("user: %+v\n\n", user)
}
