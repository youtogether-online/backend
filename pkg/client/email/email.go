package email

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"io"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type MailSender struct {
	c   *smtp.Client
	cfg *conf.Config
}

func NewEmailSender(c *smtp.Client, cfg *conf.Config) *MailSender {
	if c == nil {
		return nil
	}
	return &MailSender{c: c, cfg: cfg}
}

func Dial(username, password, host string, port int) *smtp.Client {

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
		time.Sleep(time.Minute * 5)
		if err = c.Noop(); err != nil {
			log.WithErr(err).Err("CLIENT CONNECTION LOST")
			return
		}
	}()

	return c
}

func (m *MailSender) Send(subj, body string, to ...string) error {
	if err := m.c.Mail(m.cfg.Email.User); err != nil {
		if err == io.EOF {
			if c := Dial(m.cfg.Email.User, m.cfg.Email.Password, m.cfg.Email.Host, m.cfg.Email.Port); c != nil {
				fmt.Println("RESET CLIENT")
				m.c = c
				return m.Send(subj, body, to...)
			}
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

	_, err = w.Write([]byte("From: " + m.cfg.Email.User + "\r\n" +
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
