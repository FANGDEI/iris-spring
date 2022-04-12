package main

import (
	"iris-jwt/config"
	"iris-jwt/router"
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	router.Router(app)

	err := app.Run(iris.Addr(config.Configuration.Port))
	if err != nil {
		log.Fatalln("failed to run the app")
	}
}
