package main

import (
	"go-demos/projects/golang-standard-lib-rest-api/controllers"
	"go-demos/projects/golang-standard-lib-rest-api/routes"
	"go-demos/projects/golang-standard-lib-rest-api/utils/caching"
	"go-demos/projects/golang-standard-lib-rest-api/utils/database"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal(err)
	}
	cache := &caching.Redis{
		Client: caching.Connect(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), 0),
	}

	userControllers := controllers.NewUserController(db, cache)
	jobControllers := controllers.NewJobController(db, cache)

	mux := http.NewServeMux()
	routes.CreateRoutes(mux, userControllers, jobControllers)

	log.Println("server start")
	if err := http.ListenAndServe("127.0.0.1:8000", mux); err != nil {
		log.Fatal(err)
	}
}
