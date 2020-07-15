package main

import "fmt"

// Transaction ...
func Transaction() {
	// 事务
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Exec("select * from doctor_tb where id=?", 7)
	if err != nil {
		session.Rollback()
		return
	}

	_, err = session.Exec("select * from user_tb where id=?", 67)
	if err != nil {
		session.Rollback()
		return
	}

	err = session.Commit()
	if err != nil {
		return
	}
	fmt.Println("transaction exec success.")
}
