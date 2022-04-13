package utils

import (
	"iris-jwt/model"

	"github.com/kataras/iris/v12"
)

func ResultWithData(ctx iris.Context, succeed bool, msg string, data interface{}) {
	ctx.JSON(model.Result{
		Succeed: succeed,
		Msg:     msg,
		Result:  data,
	})
}

func ResultWithoutData(ctx iris.Context, succeed bool, msg string) {
	ctx.JSON(model.Result{
		Succeed: succeed,
		Msg:     msg,
	})
}
