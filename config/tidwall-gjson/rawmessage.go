package main

import (
	"encoding/json"
	"fmt"
)

type UpLoadSomething struct {
	Type   string
	Object interface{}
}

type File struct {
	FileName string
}

type Png struct {
	Wide  int
	Hight int
}

func rawmessage() {
	input := `
    {
        "type": "File",
        "object": {
            "filename": "for test"
        }
    }
    `
	var object json.RawMessage
	ss := UpLoadSomething{
		Object: &object,
	}
	fmt.Println(ss)
	if err := json.Unmarshal([]byte(input), &ss); err != nil {
		panic(err)
	}
	switch ss.Type {
	case "File":
		var f File
		if err := json.Unmarshal(object, &f); err != nil {
			panic(err)
		}
		println(f.FileName)
	case "Png":
		var p Png
		if err := json.Unmarshal(object, &p); err != nil {
			panic(err)
		}
		println(p.Wide)
	}
}
