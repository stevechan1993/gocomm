package goemail

import (
	"testing"
)

const (
	testTo1  = ""
	testTo2  = "to2@example.com"
	testFrom = ""
	testpwd  = ""
	testBody = "Test message"
	testMsg  = "To: " + testTo1 + ", " + testTo2 + "\r\n" +
		"From: " + testFrom + "\r\n" +
		"Mime-Version: 1.0\r\n" +
		"Date: Wed, 25 Jun 2014 17:46:00 +0000\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"Content-Transfer-Encoding: quoted-printable\r\n" +
		"\r\n" +
		testBody
)

func TestSend(T *testing.T) {
	m := &EmailMessage{
		FromEmail:       testFrom,
		Toers:           []string{testTo1},
		Subject:         "测试邮件",
		BodyContentType: "text/html",
		Body:            []byte(testBody),
	}
	auth := NewLoginAuth(testFrom, testpwd, "smtp.163.com")
	err := Send("smtp.163.com:25", auth, m)
	if err != nil {
		//T.Error(err)
	}
	// mail.Message{}
}
