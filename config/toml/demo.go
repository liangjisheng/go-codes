package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

func demo() {
	data, err := ioutil.ReadFile("demo.toml")
	if err != nil {
		log.Println("read demo.toml err: ", err)
		return
	}

	var conf Config
	if _, err := toml.Decode(string(data), &conf); err != nil {
		log.Println("decode config err: ", err)
		return
	}

	fmt.Printf("%+v\n", conf)
}
