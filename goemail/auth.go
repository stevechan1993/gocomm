package goemail

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
)

//loginAuth实现smtp.Auth的接口，实现登录身份验证机制的验证.
//扩展net/smtp中实现的认证机制，net/smtp中仅实现了PlainAuth和CRAMMD5Auth
type loginAuth struct {
	username string
	password string
	host     string
}

var _ smtp.Auth = &loginAuth{}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "LOGIN" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("mail: unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, errors.New("mail: wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}

	switch {
	case bytes.EqualFold(fromServer, []byte("Username:")):
		return []byte(a.username), nil
	case bytes.EqualFold(fromServer, []byte("Password:")):
		return []byte(a.password), nil
	default:
		return nil, fmt.Errorf("mail: unexpected server challenge: %s", fromServer)
	}
}
