package handlers

import (
	"iris-jwt/model"
	"log"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx iris.Context) {
	var userDto UserDto
	err := ctx.ReadJSON(&userDto)
	if err != nil {
		log.Println("Login params is wrong")
	}
	user := model.SelectUserInformationByUsername(userDto.Username)
	if user.Password == userDto.Password {
		m := make(map[string]string)
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Id":         user.Id,
			"Username":   user.Username,
			"Permission": user.Permission,
			"iss":        "FANG",
			"iat":        time.Now().Unix(),
			// 设置三十天后token过期
			"exp": time.Now().Add(time.Minute * 60 * 24 * 30).Unix(),
		})
		tokenString, _ := token.SignedString([]byte("dyw is a big SB"))
		m["token"] = tokenString
		ctx.JSON(Result{
			Succeed: true,
			Msg:     "登录成功",
			Result:  m,
		})
		// redis key 设置为 30 天后过期
		model.SetValueWithExpire("Bearer "+tokenString, "token value, insignificance", time.Minute*60*24*30)
		return
	}
	ctx.JSON(Result{
		Succeed: false,
		Msg: "用户名或密码错误",
	})
}

func Register(ctx iris.Context) {

}

func Service(ctx iris.Context) {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims
	ctx.JSON(jwtInfo)
	ctx.HTML("<h1> this is a service </h1>")
}
