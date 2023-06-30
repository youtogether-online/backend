package email

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"io"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type MailClient struct {
	c        *smtp.Client
	host     string
	port     int
	username string
	password string
}

func NewMailClient(host string, port int, username string, password string) *MailClient {
	m := &MailClient{host: host, port: port, username: username, password: password}
	if err := m.dial(); err != nil {
		log.WithErr(err).Warn("can't open email connection")
		return nil
	}
	return m
}

func (m *MailClient) dial() error {
	if err := m.Close(); err != nil {
		return fmt.Errorf("can't close email connection: %v", err)
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", m.host, m.port), 10*time.Second)
	if err != nil {
		return fmt.Errorf("can't find specified email by HOST:PORT. Maybe HOST and PORT aren't correct?: %v", err)
	}

	SSL := m.port == 465
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	if SSL {
		conn = tls.Client(conn, cfg)
	}

	c, err := smtp.NewClient(conn, m.host)
	if err != nil {
		return fmt.Errorf("can't create client from email connection: %v", err)
	}

	if err = c.Hello("localhost"); err != nil {
		return fmt.Errorf("can't HELLO to specified email: %v", err)
	}

	if !SSL {
		if ok, _ := c.Extension("STARTTLS"); ok {
			if err = c.StartTLS(cfg); err != nil {
				if err = c.Close(); err != nil {
					return fmt.Errorf("can't HELLO to specified email: %v", err)
				}
				return fmt.Errorf("can't start email TLS connection: %v", err)
			}
		}
	}

	var auth smtp.Auth

	if m.username != "" {
		if ok, auths := c.Extension("AUTH"); ok {
			if strings.Contains(auths, "CRAM-MD5") {
				auth = smtp.CRAMMD5Auth(m.username, m.password)
			} else if strings.Contains(auths, "LOGIN") &&
				!strings.Contains(auths, "PLAIN") {
				auth = newLoginAuth(m.username, m.password, m.host)
			} else {
				auth = smtp.PlainAuth("", m.username, m.password, m.host)
			}
		}
	}

	if err = c.Auth(auth); err != nil {
		if err = c.Close(); err != nil {
			return fmt.Errorf("can't close email connection: %v", err)
		}
		return fmt.Errorf("can't authorize specified USER email. Maybe USERNAME and PASSWORD aren't correct?: %v", err)
	}

	go checkConnection(c)

	m.c = c

	return nil
}

func (m *MailClient) Send(subj, body string, to ...string) error {
	if err := m.c.Mail(m.username); err != nil {
		if err == io.EOF {
			if err = m.dial(); err != nil {
				return err
			}
			return m.Send(subj, body, to...)
		}
		return err
	}

	fmt.Println("HERE")

	for _, addr := range to {
		if err := m.c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := m.c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("From: " + m.username + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subj + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		base64.StdEncoding.EncodeToString([]byte(body))),
	)
	fmt.Println("WROTE")
	if err != nil {
		_ = w.Close()
		return err
	}
	return w.Close()
}

func (m *MailClient) Close() error {
	if m.c != nil {
		return m.c.Quit()
	}
	return nil
}

func checkConnection(c *smtp.Client) {
	time.Sleep(time.Second * 10)
	for {
		if err := c.Noop(); err != nil {
			log.WithErr(err).Err("CLIENT CONNECTION LOST")
			break
		}
		time.Sleep(time.Minute * 5)
	}
}
