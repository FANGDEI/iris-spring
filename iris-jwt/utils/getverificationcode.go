package utils

import (
	"math/rand"
	"time"
)

const BASE_VALUE = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetVerificationCode() string {
	code := make([]byte, 6)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 6; i++ {
		code = append(code, BASE_VALUE[rand.Intn(len(BASE_VALUE))])
	}
	return string(code)
}
