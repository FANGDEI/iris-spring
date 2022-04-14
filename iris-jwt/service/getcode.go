package service

import (
	"context"
	"iris-jwt/proto/getcode"
	"iris-jwt/utils"
	"log"
	"regexp"
	"time"
)

type GetCodeServer struct {
}

func (s *GetCodeServer) GetCode(ctx context.Context, request *getcode.CodeRequest) (*getcode.CodeResponse, error) {
	return getCode(request.GetEmail()), nil
}

func getCode(email string) *getcode.CodeResponse {
	// 正则判断 [1-9][0-9]+@qq.com
	matched, _ := regexp.MatchString(`[1-9][0-9]+@qq.com`, email)
	if !matched {
		return &getcode.CodeResponse{Succeed: false, Message: "请输入正确的邮箱"}
	}

	// 获取验证码
	code := utils.GetVerificationCode()

	// 发送邮件验证码
	err := utils.SendEmail("Fang&Devil", email, "验证码", "验证码十五分钟之内有效<br>验证码: "+code)
	if err != nil {
		log.Println("failed to send the eamil.", err)
		return &getcode.CodeResponse{Succeed: false, Message: "未知原因发送验证码失败"}
	}

	// code 存入 redis 十五分钟后失效
	err = utils.SetValueWithExpire(code, email, time.Minute*15)
	if err != nil {
		log.Println("failed to put verifycode into redis.", err)
	}

	return &getcode.CodeResponse{Succeed: true, Message: "验证码已发送"}
}
