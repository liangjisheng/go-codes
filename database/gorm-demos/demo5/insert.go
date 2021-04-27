package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

// 生成随机的字符串
func gainRandomString(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for v := range b {
		b[v] = s[rand.Intn(len(s))]
	}
	return string(b)
}

// 生成md5字符串
func gainRandomMd5String(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func gainData(d chan Admin) {
	for i := 0; i < 20; i++ {
		name := gainRandomString(9)
		pwd := gainRandomMd5String(name)
		data := Admin{User: name, Password: pwd}
		d <- data
	}
	close(d)
}

func insert() {
	datas := make(chan Admin, 20)
	go gainData(datas)

	for v := range datas {
		// NewRecord check if value's primary key is blank
		db.NewRecord(v)
		// Create insert the value into database
		db.Create(&v)
	}
}
