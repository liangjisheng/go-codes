package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func demo1() {
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/05blog?charset=utf8")
	if err != nil {
		fmt.Println("open db fail", err)
	}

	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()

	// selectDemo(db)
	// insertUpdateDelete(db)
	transaction(db)
}

func selectDemo(db *sql.DB) {
	//查询数据，指定字段名，返回sql.Rows结果集
	rows, err := db.Query("select id, name from user")
	if err != nil {
		fmt.Println(err)
	}
	id := 0
	name := ""
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

	// 查询数据，取所有字段
	rows2, _ := db.Query("select * from user")
	// 返回所有列
	cols, _ := rows2.Columns()
	fmt.Println(cols)
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		// 填充数据
		rows2.Scan(scans...)
		// 每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	fmt.Println(result)

	// 预处理语句
	stmt, _ := db.Prepare("select id, name from user where id = ?")
	rows4, _ := stmt.Query(1)
	// 注意这里需要Next()下，不然下面取不到值
	rows4.Next()
	rows4.Scan(&id, &name)
	fmt.Println(id, name)
}

func insertUpdateDelete(db *sql.DB) {
	ret, _ := db.Exec("insert into user (name, address, phone) values ('jack', 'NewYork', '123')")
	// 获取插入ID
	insID, _ := ret.LastInsertId()
	fmt.Println("insID:", insID)

	ret2, _ := db.Exec("update user set name = '000' where id = ?", insID)
	// 获取影响行数
	affNum, _ := ret2.RowsAffected()
	fmt.Println("affect num:", affNum)

	ret3, _ := db.Exec("delete from user where id = ?", insID)
	// 获取影响行数
	delNum, _ := ret3.RowsAffected()
	fmt.Println("delete num:", delNum)
}

func transaction(db *sql.DB) {
	tx, _ := db.Begin()

	id := 1
	ret, _ := tx.Exec("update user set name = 'alice' where id = ?", id)
	ret1, _ := tx.Exec("update user set address = 'wenshui' where id = ?", id)
	updNum, _ := ret.RowsAffected()
	updNum1, _ := ret1.RowsAffected()
	fmt.Println("updNum:", updNum)
	fmt.Println("updNum1:", updNum1)

	if updNum > 0 && updNum1 > 0 {
		fmt.Println("commit")
		tx.Commit() // 只有两条更新同时成功，那么才提交
	} else {
		fmt.Println("rollback")
		tx.Rollback() //否则回滚
	}
}
