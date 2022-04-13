package middleware

import (
	"iris-jwt/constant"
	"iris-jwt/model"
	"iris-jwt/utils"
	"log"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// 生成验证token的中间件
func JWTMiddleware() *jwt.Middleware {
	// 前端发送token时带上前缀 Bearer + 空格
	j := jwt.New(jwt.Config{
		// 从请求头中的Authorization获取token进行验证
		Extractor: jwt.FromAuthHeader,
		// 密钥
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.TOKEN_SECRET), nil
		},
		// 签名算法
		SigningMethod: jwt.SigningMethodHS256,
		// 错误处理函数
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil {
				return
			}
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(model.Result{
				Succeed: false,
				Msg:     err.Error(),
			})
		},
	})
	return j
}

// Redis验证 token 中间件
func JWTRedisVerifyHandler(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")
	_, err := utils.GetValue(token)
	if err != nil || token == "" {
		log.Println("invalid token or the token is nil")
		ctx.JSON(model.Result{
			Succeed: false,
			Msg:     "the token is invalid or the token is null",
		})
		return
	}
	ctx.Next()
}
