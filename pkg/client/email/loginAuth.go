package email

import (
	"bytes"
	"fmt"
	"net/smtp"
)

type loginAuth struct {
	username string
	password string
	host     string
}

func newLoginAuth(username string, password string, host string) *loginAuth {
	return &loginAuth{username: username, password: password, host: host}
}

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
			return "", nil, fmt.Errorf("unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, fmt.Errorf("wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}

	switch {
	case bytes.Equal(fromServer, []byte("Username:")):
		return []byte(a.username), nil
	case bytes.Equal(fromServer, []byte("Password:")):
		return []byte(a.password), nil
	default:
		return nil, fmt.Errorf("unexpected server challenge: %s", fromServer)
	}
}
