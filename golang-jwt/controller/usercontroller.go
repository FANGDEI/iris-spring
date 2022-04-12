package controller

import (
	"golang-jwt/model"
	"golang-jwt/model/dto"
	"golang-jwt/service"
	"log"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: service.NewUserService()}
}

func (g *UserController) PostLogin(ctx iris.Context) model.Result {
	var userDto dto.UserDto
	err := ctx.ReadJSON(&userDto)
	if err != nil {
		log.Println("failed to read the params")
		return model.Result{
			Succeed: false,
			Msg:     "未知错误, 请联系管理员",
		}
	}
	return g.Service.Login(userDto)
}
