package main

import (
	"iris-jwt/config"
	"iris-jwt/core"
	"iris-jwt/router"
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	core.ConsulFindServer("cloud-grpc-server-01")

	app := iris.New()

	router.Router(app)

	err := app.Run(iris.Addr(config.Configuration.Port))
	if err != nil {
		log.Fatalln("failed to run the app")
	}
}
