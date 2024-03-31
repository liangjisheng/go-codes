package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func unmarshal() {
	data := `[{"Level":"debug","Msg":"File: \"test.txt\" Not Found"},` +
		`{"Level":"","Msg":"Logic error"}]`

	var dbgInfos []map[string]string
	json.Unmarshal([]byte(data), &dbgInfos)
	fmt.Println(dbgInfos)
}

type configuration struct {
	Enabled bool   `json:"enabled"`
	Path    string `json:"path"`
	Number  int64  `json:"number"`
}

func unmarshal1() {
	// file, _ := os.Open("conf.json")
	// defer file.Close()

	// decoder := json.NewDecoder(file)
	// conf := configuration{}
	// err := decoder.Decode(&conf)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(conf.Path)

	content, _ := ioutil.ReadFile("conf.json")
	conf := configuration{}
	json.Unmarshal(content, &conf)
	fmt.Printf("%+v\n", conf)
}
