package main

import (
	"flag"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var (
	rps int
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[GIN] ")
	log.SetOutput(gin.DefaultWriter)

	flag.IntVar(&rps, "rps", 100, "request per second")
}

func leakBucket(rps int) gin.HandlerFunc {
	limit := ratelimit.New(rps)
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

func main() {
	flag.Parse()

	app := gin.Default()
	app.Use(leakBucket(rps))

	app.GET("/rate", func(ctx *gin.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
	app.Run(":8080")
}
