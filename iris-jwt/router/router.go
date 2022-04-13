package router

import (
	"iris-jwt/handlers"
	"iris-jwt/middleware"

	"github.com/kataras/iris/v12"
)

func Router(app *iris.Application) {
	j := middleware.JWTMiddleware()

	app.Post("/login", handlers.Login)
	app.Post("/getCode", handlers.GetCode)
	app.Post("/register", handlers.Register)
	app.Get("/userService", j.Serve, middleware.JWTRedisVerifyHandler, handlers.UserService)
}
