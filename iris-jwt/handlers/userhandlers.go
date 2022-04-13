package handlers

import (
	"iris-jwt/constant"
	"iris-jwt/repo"
	"iris-jwt/utils"
	"log"
	"regexp"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

func Login(ctx iris.Context) {
	var userDto UserDto
	err := ctx.ReadJSON(&userDto)
	if err != nil {
		log.Println("failed to read params.", err)
		utils.ResultWithoutData(ctx, false, "未知原因登录失败")
		return
	}

	user, err := repo.SelectUserInformationByUsername(userDto.Username)
	if err != nil {
		log.Println("failed to get user information.", err)
		utils.ResultWithoutData(ctx, false, "未知原因登录失败")
		return
	}

	println(user.Password)

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password)); err == nil {
		m := make(map[string]string)
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":         user.Id,
			"username":   user.Username,
			"permission": user.Permission,
			"iss":        "FANG",
			"iat":        time.Now().Unix(),
			// 设置三十天后token过期
			"exp": time.Now().Add(time.Minute * 60 * 24 * 30).Unix(),
		})
		tokenString, _ := token.SignedString([]byte(constant.TOKEN_SECRET))
		m["token"] = tokenString

		utils.ResultWithData(ctx, true, "登录成功", m)

		// redis key 设置为 30 天后过期
		utils.SetValueWithExpire("Bearer "+tokenString, "token value, insignificance", time.Minute*60*24*30)
		return
	}

	utils.ResultWithoutData(ctx, false, "用户名或密码错误")
}

func GetCode(ctx iris.Context) {
	var userDto UserDto
	ctx.ReadJSON(&userDto)

	// 正则判断 [1-9][0-9]+@qq.com
	matched, _ := regexp.MatchString(`[1-9][0-9]+@qq.com`, userDto.Email)
	if !matched {
		utils.ResultWithoutData(ctx, false, "请输入正确的邮箱")
		return
	}

	// 获取验证码
	code := utils.GetVerificationCode()

	// 发送邮件验证码
	err := utils.SendEmail("Iris-Spring", userDto.Email, "验证码", "验证码十五分钟之内有效<br>验证码: "+code)
	if err != nil {
		log.Println("failed to send the eamil.", err)
		utils.ResultWithoutData(ctx, false, "未知原因验证码发送错误")
		return
	}

	// code 存入 redis 十五分钟后失效
	err = utils.SetValueWithExpire(code, userDto.Email, time.Minute*15)
	if err != nil {
		log.Println("failed to put verifycode into redis.", err)
	}

	utils.ResultWithoutData(ctx, true, "验证码已发送")
}

func Register(ctx iris.Context) {
	var userDto UserDto
	err := ctx.ReadJSON(&userDto)
	if err != nil {
		log.Println("failed to read params.", err)
		utils.ResultWithoutData(ctx, false, "未知原因注册失败")
		return
	}

	// 用户名不能为空
	if userDto.Username == "" {
		utils.ResultWithoutData(ctx, false, "用户名不能为空")
		return
	}

	// 密码长度不能小于6位
	if len(userDto.Password) < 6 {
		utils.ResultWithoutData(ctx, false, "密码不能小于六位，请重新输入密码")
		return
	}

	email, errGet := utils.GetValue(userDto.Code)
	if errGet != nil {
		log.Println("failed to get value from redis.", err)
		utils.ResultWithoutData(ctx, false, "未知原因注册失败")
		return
	}

	// 验证码错误
	if email != userDto.Email {
		utils.ResultWithoutData(ctx, false, "验证码错误或已过期，请检查")
		return
	}

	// redis 中删除 code
	errDeleteValue := utils.DeleteValue(userDto.Code)
	if errDeleteValue != nil {
		log.Println("failed to delete value from redis")
		utils.ResultWithoutData(ctx, false, "未知原因注册失败")
		return
	}

	// 密码加密
	hash, errHash := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Println("failed to get the hash with the password")
		utils.ResultWithoutData(ctx, false, "未知原因注册失败")
	}
	encodePassword := string(hash)

	cnt, errInsert := repo.InsertUser(userDto.Username, encodePassword, userDto.Email)
	if errInsert != nil || cnt == 0 {
		log.Println("failed to insert user into db.", err)
		utils.ResultWithoutData(ctx, false, "未知原因注册失败")
		return
	}

	utils.ResultWithoutData(ctx, true, "注册成功")
}

func UserService(ctx iris.Context) {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims
	ctx.JSON(jwtInfo)
	ctx.HTML("<h1> this is a service </h1>")
}

func Logout(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")

	err := utils.DeleteValue(token)
	if err != nil {
		log.Println("failed to delete token")
		utils.ResultWithoutData(ctx, false, "未知原因退出登录失败")
		return
	}

	utils.ResultWithoutData(ctx, true, "退出登录成功")
}
