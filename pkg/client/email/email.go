package email

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type MailClient struct {
	host     string
	port     int
	username string
	password string
	SSL      bool
	authType authType
	tlsConf  *tls.Config
}

type authType string

const (
	crammd5 authType = "CRAM-MD5"
	login   authType = "LOGIN"
	plain   authType = "PLAIN"
)

func NewMailClient(host string, port int, username string, password string) *MailClient {
	m := &MailClient{host: host, port: port, username: username, password: password,
		SSL: port == 465, tlsConf: &tls.Config{InsecureSkipVerify: true}}

	client, err := m.dial()
	if err != nil {
		log.WithErr(err).Warn("can't open email connection")
	} else if err = client.Quit(); err != nil {
		log.WithErr(err).Warn("can't close email connection")
	}
	return m
}

func (m *MailClient) dial() (*smtp.Client, error) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", m.host, m.port), 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("can't find specified email by HOST:PORT. Maybe HOST and PORT aren't correct?: %v", err)
	}

	if m.SSL {
		conn = tls.Client(conn, m.tlsConf)
	}

	c, err := smtp.NewClient(conn, m.host)
	if err != nil {
		return nil, err
	}

	if err = c.Hello("localhost"); err != nil {
		return nil, err
	}

	if !m.SSL {
		if ok, _ := c.Extension("STARTTLS"); ok {
			if err = c.StartTLS(m.tlsConf); err != nil {
				if err = c.Close(); err != nil {
					return nil, err
				}
				return nil, fmt.Errorf("can't start email TLS connection: %v", err)
			}
		}
	}

	var auth smtp.Auth

	if m.username != "" {
		if ok, auths := c.Extension("AUTH"); ok {
			if strings.Contains(auths, "CRAM-MD5") {
				auth, m.authType = smtp.CRAMMD5Auth(m.username, m.password), crammd5
			} else if strings.Contains(auths, "LOGIN") &&
				!strings.Contains(auths, "PLAIN") {
				auth, m.authType = newLoginAuth(m.username, m.password, m.host), login
			} else {
				auth, m.authType = smtp.PlainAuth("", m.username, m.password, m.host), plain
			}
		}
	}

	if err = c.Auth(auth); err != nil {
		if err = c.Close(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("can't log in to email. Maybe USERNAME and PASSWORD aren't correct?: %v", err)
	}

	return c, nil
}

func (m *MailClient) fasterDial() (*smtp.Client, error) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", m.host, m.port), 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("can't find specified email by HOST:PORT. Maybe HOST and PORT aren't correct?: %v", err)
	}

	if m.SSL {
		conn = tls.Client(conn, m.tlsConf)
	}

	c, err := smtp.NewClient(conn, m.host)
	if err != nil {
		return nil, err
	}

	if err = c.Hello("localhost"); err != nil {
		return nil, err
	}

	if !m.SSL {
		if ok, _ := c.Extension("STARTTLS"); ok {
			if err = c.StartTLS(m.tlsConf); err != nil {
				if err = c.Close(); err != nil {
					return nil, err
				}
				return nil, fmt.Errorf("can't start email TLS connection: %v", err)
			}
		}
	}

	var auth smtp.Auth

	switch m.authType {
	case crammd5:
		auth = smtp.CRAMMD5Auth(m.username, m.password)
	case login:
		auth = newLoginAuth(m.username, m.password, m.host)
	case plain:
		auth = smtp.PlainAuth("", m.username, m.password, m.host)
	}

	if err = c.Auth(auth); err != nil {
		if err = c.Close(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("can't log in to email. Maybe USERNAME and PASSWORD changed?: %v", err)
	}

	return c, nil
}

func (m *MailClient) DialAndSend(subj, body string, to ...string) error {
	c, err := m.fasterDial()
	if err != nil {
		return err
	}

	if err = c.Mail(m.username); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, werr := w.Write([]byte("From: " + m.username + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subj + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		base64.StdEncoding.EncodeToString([]byte(body))),
	)
	if err = w.Close(); err != nil {
		return err
	}

	if werr != nil {
		return err
	}

	return c.Quit()
}
