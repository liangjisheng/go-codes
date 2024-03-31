package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type person struct {
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bsin:"phone"`
}

func insert() {
	session, err := mgo.Dial("localhost:27017")
	checkErr(err)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior
	session.SetMode(mgo.Monotonic, true)

	// 切换到test数据库 alice 集合
	c := session.DB("test").C("alice")

	// insert
	err = c.Insert(&person{"alice", "123"},
		&person{"lisi", "456"})
	checkErr(err)

	results := []*person{}
	err = c.Find(bson.M{"name": "alice"}).All(&results)
	checkErr(err)

	if len(results) > 0 {
		for i, result := range results {
			fmt.Printf("data:%d name:%s phone:%s\n", i, result.Name, result.Phone)
		}
	} else {
		fmt.Println("not find data")
	}
}
