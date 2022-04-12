package main

import (
	"iris-jwt/router"
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	router.Router(app)

	err := app.Run(iris.Addr(":7777"))
	if err != nil {
		log.Fatalln("failed to run the app")
	}
}
