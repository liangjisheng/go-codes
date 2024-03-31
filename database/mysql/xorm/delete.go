package main

import (
	"fmt"
	"xorm/xorm-models/models"
)

// Delete ...
func Delete() {
	// 删除
	docDel := models.DoctorTb{Name: "doctor-alice3"}
	iDel, _ := engine.Delete(&docDel)
	fmt.Println("删除结果：", iDel)
	fmt.Println()

	// 删除方式2
	iDel1, _ := engine.Exec("delete from doctor_tb where id = ?", 14)
	fmt.Println("删除结果：", iDel1)
	fmt.Println()
}
