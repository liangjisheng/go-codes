package main

import (
	"gofeat/go-build/app"
	"log"
)

var (
	Version   string
	BuildUser string
)

func main() {
	log.Println("Version", Version)
	log.Println("BuildUser", BuildUser)

	app.Vars()
}
