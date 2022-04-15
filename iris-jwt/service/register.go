package service

import (
	"context"
	re "iris-jwt/proto/register"
	"iris-jwt/repo"
	"iris-jwt/utils"
	"log"
	"net"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RegisterServer struct {
}

func (s *RegisterServer) Register(ctx context.Context, request *re.RegisterRequest) (*re.RegisterResponse, error) {
	return register(request.GetUsername(), request.GetPassword(), request.GetCode(), request.GetEmail()), nil
}

func register(username, password, code, email string) *re.RegisterResponse {
	// 用户名不能为空
	if username == "" {
		return &re.RegisterResponse{Succeed: false, Message: "用户名不能为空"}
	}

	// 密码长度不能小于6位
	if len(password) < 6 {
		return &re.RegisterResponse{Succeed: false, Message: "密码不能小于六位，请重新输入密码"}
	}

	redisEmail, errGet := utils.GetValue(code)
	if errGet != nil {
		log.Println("failed to get value from redis.", errGet)
		return &re.RegisterResponse{Succeed: false, Message: "未知原因注册失败"}
	}

	// 验证码错误
	if redisEmail != email {
		return &re.RegisterResponse{Succeed: false, Message: "验证码错误或已过期，请检查"}
	}

	// redis 中删除 code
	errDeleteValue := utils.DeleteValue(code)
	if errDeleteValue != nil {
		log.Println("failed to delete value from redis")
		return &re.RegisterResponse{Succeed: false, Message: "未知原因注册失败"}
	}

	// 密码加密
	hash, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Println("failed to get the hash with the password.")
		return &re.RegisterResponse{Succeed: false, Message: "未知原因注册失败"}
	}
	encodePassword := string(hash)

	cnt, errInsert := repo.InsertUser(username, encodePassword, email)
	if errInsert != nil || cnt == 0 {
		log.Println("failed to insert user into db.", errInsert)
		return &re.RegisterResponse{Succeed: false, Message: "未知原因注册失败"}
	}

	return &re.RegisterResponse{Succeed: true, Message: "注册成功"}
}

func RunRegisterService() {
	log.Println("Register service is running...")
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Println("failed to listen.", err)
	}

	s := grpc.NewServer()
	re.RegisterRegisterServer(s, &RegisterServer{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Println("Register failed to server.", err)
	}
}
