package email

import (
	"crypto/tls"
	"fmt"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net/smtp"
	"strings"
)

// Open smtp connection, start TLS and authorize the user
func Open(username, password, host string, port int) *smtp.Client {

	c, err := smtp.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.WithErr(err).Warn("can't connect to specified email HOST:PORT")
		return nil
	}

	if err = c.Hello("localhost"); err != nil {
		return nil
	}

	// only no SSL
	if ok, _ := c.Extension("STARTTLS"); ok {
		if err = c.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
			c.Close()
			log.WithErr(err).Warn("can't start email TLS connection")
			return nil
		}
	}

	var auth smtp.Auth

	if username != "" {
		if ok, auths := c.Extension("AUTH"); ok {
			if strings.Contains(auths, "CRAM-MD5") {
				auth = smtp.CRAMMD5Auth(username, password)
			} else {
				auth = smtp.PlainAuth("", username, password, host)
			}
		}
	}

	if err = c.Auth(auth); err != nil {
		c.Close()
		log.WithErr(err).Warn("can't authorize specified USER email")
		return nil
	}

	return c
}
