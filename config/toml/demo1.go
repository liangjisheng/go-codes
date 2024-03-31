package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type song struct {
	Name     string
	Duration duration
}

type songs struct {
	Song []song
}

func demo1() {
	data, err := ioutil.ReadFile("demo1.toml")
	if err != nil {
		log.Println("read demo.toml err: ", err)
		return
	}

	var favorites songs
	if _, err := toml.Decode(string(data), &favorites); err != nil {
		log.Println("decode config err: ", err)
		return
	}

	for _, s := range favorites.Song {
		fmt.Printf("%s (%s)\n", s.Name, s.Duration)
	}
}
