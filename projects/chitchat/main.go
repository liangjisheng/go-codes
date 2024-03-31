package main

import (
	"log"
	"net/http"

	"chitchat/config"
	"chitchat/routes"
)

func main() {
	// config := config.LoadConfig()
	r := routes.NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir(config.ViperConfig.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + config.ViperConfig.App.Address)
	err := http.ListenAndServe(config.ViperConfig.App.Address, nil)
	if err != nil {
		log.Println("An error occured starting HTTP listener at " + config.ViperConfig.App.Address)
		log.Println("Error: " + err.Error())
	}
}
