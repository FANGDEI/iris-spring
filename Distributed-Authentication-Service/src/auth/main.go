package main

import (
	_ "distributed/authentication/store/local"
	"flag"

	"github.com/kataras/iris/v12"
)

var (
	port = flag.String("port", "7777", "Port iris to listen")
)

func main() {
	app := iris.New()

	app.Run(iris.Addr("0.0.0.0:" + *port))
}
