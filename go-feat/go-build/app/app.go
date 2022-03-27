package app

import "log"

var (
	BuildUser string
	BuildTime string
)

func Vars() {
	log.Println("app.BuildUser", BuildUser)
	log.Println("app.BuildTime", BuildTime)
}
