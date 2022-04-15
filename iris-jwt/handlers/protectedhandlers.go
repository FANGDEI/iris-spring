package handlers

import "github.com/kataras/iris/v12"

func AdminService(ctx iris.Context) {
	ctx.HTML("<h1> this is a service for admin </h1>")
}
