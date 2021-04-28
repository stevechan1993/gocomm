package goemail

import (
	"net/smtp"
)

// ServerHost      string //邮箱服务器地址
// ServerPort      string   //邮箱服务器的端口
// FromPasswd      string //密码

//NewPlainAuth
func NewPlainAuth(fromEmail string, pwd string, host string) smtp.Auth {
	return smtp.PlainAuth("", fromEmail, pwd, host)
}

func NewLoginAuth(fromEmail string, pwd string, host string) smtp.Auth {
	return &loginAuth{
		username: fromEmail,
		password: pwd,
		host:     host,
	}
}

func Send(addr string, auth smtp.Auth, m *EmailMessage) error {
	return smtp.SendMail(addr, auth, m.FromEmail, m.Toers, m.Bytes())
}
