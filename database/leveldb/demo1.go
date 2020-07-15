package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

const ID_FIELD = "id"

func initdb() {
	var err error
	db, err = leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}

	_, err = db.Get([]byte(ID_FIELD), nil)
	if err != nil {
		db.Put([]byte(ID_FIELD), []byte("10000"), nil)
	}
}

func getNextID() int {
	ids, err := db.Get([]byte(ID_FIELD), nil)
	if err != nil {
		fmt.Println(err)
	}

	id := byte2int(ids)
	db.Put([]byte(ID_FIELD), int2Byte(id+1), nil)
	return id
}

func byte2int(val []byte) int {
	var result int
	result, _ = strconv.Atoi(string(val))
	return result
}

func int2Byte(val int) []byte {
	result := []byte(strconv.Itoa(val))
	return result
}

func demo1() {
	initdb()

	var val int
	val = getNextID()
	fmt.Println(val)
	val = getNextID()
	fmt.Println(val)

	db.Close()
}
