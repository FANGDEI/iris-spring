package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func AuthorizationMiddleware(ctx iris.Context) {
	claimsMap := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	permission := claimsMap["permission"]
	if int(permission.(float64)) < 1 {
		ctx.HTML("<h1> You don't have permission to visit </h1>")
		return
	}
	ctx.HTML("<h1> this is a service for admin </h1>")
}
