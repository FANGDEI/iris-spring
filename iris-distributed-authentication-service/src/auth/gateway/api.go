package gateway

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func New() *iris.Application {
	app := iris.New()

	// cors
	app.UseRouter(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	}))

	return app
}
