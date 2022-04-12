package utils

import (
	"iris-jwt/config"
	"log"

	"github.com/go-gomail/gomail"
)

var m *gomail.Message

func SendEmail(sender, receiver, subject, body string) {
	m = gomail.NewMessage()
	// 收件人
	m.SetHeader("To", receiver)
	// 发件人 第三个参数可设置发件人别名
	m.SetAddressHeader("From", config.Configuration.Mail.FromEamil, sender)
	// 主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(config.Configuration.Mail.ServerHost, config.Configuration.Mail.ServerPort, config.Configuration.Mail.FromEamil, config.Configuration.Mail.Password)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		log.Println("发送邮件失败")
		return
	}
}
