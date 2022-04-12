package service

import (
	"golang-jwt/model"
	"golang-jwt/model/dto"
)

type UserService interface {
	Login(userDto dto.UserDto) model.Result
	GetInfo() model.Result
	GetCode(email string) model.Result
	Register(userDto dto.UserDto) model.Result
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (user userService) Login(userDto dto.UserDto) model.Result {
	return model.Result{
		Succeed: true,
		Msg:     "登录成功",
	}
}

func (user userService) GetInfo() model.Result {
	return model.Result{}
}

func (user userService) GetCode(email string) model.Result {
	return model.Result{}
}

func (user userService) Register(userDto dto.UserDto) model.Result {
	return model.Result{}
}
