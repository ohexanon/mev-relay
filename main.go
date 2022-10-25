package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cornelk/hashmap"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {

	requestMap := hashmap.New[string, []time.Time]()

	app := iris.New()
	app.Use(iris.Compression)
	app.Post("/", memoryRateLimiting(requestMap), handlePost)
	app.Get("/", memoryRateLimiting(requestMap), func(ctx iris.Context) { ctx.WriteString("ok") })

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
