package goemail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type EmailHeader struct {
	Key   string
	Value string
}

type EmailMessage struct {
	FromEmail string //发件人邮箱地址
	//	Header          []EmailHeader
	Toers           []string //邮件接收人，
	Subject         string   //主题
	BodyContentType string   //默认值text/html
	BodyCharset     string   //字符编码设定默认utf-8
	Body            []byte   //邮件正文内容
	//TODO 添加附件
}

func NewEmailMessage(fromemail string, subject string, toers ...string) *EmailMessage {
	return &EmailMessage{
		FromEmail:       fromemail,
		Toers:           toers,
		BodyContentType: "text/html",
		BodyCharset:     "utf-8",
	}
}

func (e *EmailMessage) AddToer(v string) {
	e.Toers = append(e.Toers, v)
}

func (e *EmailMessage) SetBody(v []byte) {
	e.Body = v
}

func (e *EmailMessage) SetBodyContentType(v string) {
	e.BodyContentType = v
}

func (e *EmailMessage) SetBodyCharset(v string) {
	e.BodyCharset = v
}

func (e *EmailMessage) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("From:%s \r\n", e.FromEmail))
	t := time.Now().Format(time.RFC1123Z)
	buf.WriteString(fmt.Sprintf("Date:%s \r\n", t))
	buf.WriteString(fmt.Sprintf("To:%s \r\n", strings.Join(e.Toers, ",")))
	var subject = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(e.Subject)) + "?="
	buf.WriteString("Subject: " + subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	//TODO 添加header
	boundary := "THIS_IS_BOUNDARY"
	buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n")
	buf.WriteString("\r\n--" + boundary + "\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=%s\r\n\r\n", e.BodyContentType, e.BodyCharset))
	buf.Write(e.Body)
	buf.WriteString("\r\n")
	buf.WriteString("--" + boundary + "--")
	//TODO添加附件
	return buf.Bytes()
}
