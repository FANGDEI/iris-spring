package utils

import (
	"iris-jwt/constant"
	"math/rand"
	"time"
)

func GetVerificationCode() string {
	code := make([]byte, 6)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 6; i++ {
		code = append(code, constant.CODE_SECRET[rand.Intn(len(constant.CODE_SECRET))])
	}
	return string(code)[6:]
}
