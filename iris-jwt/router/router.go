package router

import (
	"iris-jwt/handlers"
	"iris-jwt/middleware"

	"github.com/kataras/iris/v12"
)

func Router(app *iris.Application) {
	j := middleware.JWTMiddleware()

	app.Get("/check", func(ctx iris.Context) {
		ctx.WriteString("the check function is ok")
	})
	app.Post("/login", handlers.Login)
	app.Post("/getCode", handlers.GetCode)
	app.Post("/register", handlers.Register)
	app.Get("/logout", j.Serve, middleware.JWTRedisVerifyMiddleware, handlers.Logout)
	app.Get("/userService", j.Serve, middleware.JWTRedisVerifyMiddleware, handlers.UserService)
	app.Get("/adminService", j.Serve, middleware.JWTRedisVerifyMiddleware, middleware.AuthorizationMiddleware, handlers.AdminService)
}
