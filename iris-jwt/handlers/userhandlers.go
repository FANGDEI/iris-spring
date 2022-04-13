package handlers

import (
	"iris-jwt/repo"
	"iris-jwt/utils"
	"log"
	"regexp"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmailDto struct {
	Email string `json:"email"`
}

func Login(ctx iris.Context) {
	var userDto UserDto
	err := ctx.ReadJSON(&userDto)
	if err != nil {
		utils.ResultWithoutData(ctx, false, "未知原因登录失败")
		return
	}

	user := repo.SelectUserInformationByUsername(userDto.Username)
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

		utils.ResultWithData(ctx, true, "登录成功", m)

		// redis key 设置为 30 天后过期
		utils.SetValueWithExpire("Bearer "+tokenString, "token value, insignificance", time.Minute*60*24*30)
		return
	}

	utils.ResultWithoutData(ctx, false, "用户名或密码错误")
}

func GetCode(ctx iris.Context) {
	var emailDto EmailDto
	ctx.ReadJSON(&emailDto)

	// 正则判断 [1-9][0-9]+@qq.com
	matched, _ := regexp.MatchString(`[1-9][0-9]+@qq.com`, emailDto.Email)
	if !matched {
		utils.ResultWithoutData(ctx, false, "请输入正确的邮箱")
		return
	}

	code := utils.GetVerificationCode()

	err := utils.SendEmail("Iris-Spring", emailDto.Email, "验证码", "验证码十五分钟之内有效<br>验证码: "+code)
	if err != nil {
		utils.ResultWithoutData(ctx, false, "未知原因验证码发送错误")
		return
	}

	// code 存入 redis 十五分钟后失效
	err = utils.SetValueWithExpire(code, emailDto.Email, time.Minute*15)
	if err != nil {
		log.Println("验证码存入redis失败")
	}

	utils.ResultWithoutData(ctx, true, "验证码已发送")
}

func Register(ctx iris.Context) {

}

func Service(ctx iris.Context) {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims
	ctx.JSON(jwtInfo)
	ctx.HTML("<h1> this is a service </h1>")
}
