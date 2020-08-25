package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

//定义SMTPInfo结构体，用于传递发送邮箱所必须的信息
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage() //调用方法创建消息实例，方便下面设置邮件发送的必要信息
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	//通过NewDialer创建SMTP拨号实例
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m) //最后调用 DialAndSend 方法打开与SMTP服务器的连接并发送邮件
}
