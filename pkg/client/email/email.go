package email

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net"
	"net/smtp"
	"strings"
	"time"
)

// Open smtp connection, start TLS and authorize the user
func Open(username, password, host string, port int) *smtp.Client {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 10*time.Second)
	if err != nil {
		log.WithErr(err).Warn("can't find specified email by HOST:PORT. Maybe HOST and PORT aren't correct?")
		return nil
	}

	SSL := port == 465
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	if SSL {
		conn = tls.Client(conn, cfg)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.WithErr(err).Warn("can't create client from email connection")
		return nil
	}

	if err = c.Hello("localhost"); err != nil {
		log.WithErr(err).Err("can't HELLO to specified email")
		return nil
	}

	if !SSL {
		if ok, _ := c.Extension("STARTTLS"); ok {
			if err = c.StartTLS(cfg); err != nil {
				if err = c.Close(); err != nil {
					log.WithErr(err).Warn("can't close email connection")
				}
				log.WithErr(err).Warn("can't start email TLS connection")
				return nil
			}
		}
	}

	var auth smtp.Auth

	if username != "" {
		if ok, auths := c.Extension("AUTH"); ok {
			if strings.Contains(auths, "CRAM-MD5") {
				auth = smtp.CRAMMD5Auth(username, password)
			} else if strings.Contains(auths, "LOGIN") &&
				!strings.Contains(auths, "PLAIN") {
				auth = &loginAuth{
					username: username,
					password: password,
					host:     host,
				}
			} else {
				auth = smtp.PlainAuth("", username, password, host)
			}
		}
	}

	if err = c.Auth(auth); err != nil {
		if err = c.Close(); err != nil {
			log.WithErr(err).Warn("can't close email connection")
		}
		log.WithErr(err).Warn("can't authorize specified USER email. Maybe USERNAME and PASSWORD aren't correct?")
		return nil
	}

	go func() {
		for {
			if err = c.Noop(); err != nil {
				log.WithErr(err).Err("email connection lost")
				break
			}
			time.Sleep(time.Second * 10)
		}
	}()

	return c
}

type loginAuth struct {
	username string
	password string
	host     string
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
			return "", nil, errors.New("unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
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
