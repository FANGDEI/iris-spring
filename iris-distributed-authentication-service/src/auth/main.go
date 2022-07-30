package main

import (
	"distributed/authentication/gateway"
	"distributed/authentication/store/local"
	"flag"
	"github.com/kataras/iris/v12"
	"log"
)

var (
	port = flag.String("port", "7777", "Port iris to listen")
)

func main() {
	local.M.DeleteUserByUsername()
	flag.Parse()

	app := gateway.New()

	err := app.Run(iris.Addr(":" + *port))
	if err != nil {
		log.Fatalf("[IRIS SERVE ERROR] %v", err)
	}
}
