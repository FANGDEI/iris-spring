package main

import (
	"golang-jwt/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(controller.NewUserController())
	app.Run(iris.Addr(":7777"))
}
