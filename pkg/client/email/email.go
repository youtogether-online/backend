package email

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"net/smtp"
)

// Open smtp connection, start TLS and authorize the user
func Open(addr string, auth smtp.Auth) *smtp.Client {
	c, err := smtp.Dial(addr)
	if err != nil {
		logrus.WithError(err).Warn("can't connect to specified email HOST:PORT")
		return nil
	}

	if err = c.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
		logrus.WithError(err).Warn("can't start email TLS connection")
		return nil
	}

	if err = c.Auth(auth); err != nil {
		logrus.WithError(err).Warn("can't authorize specified USER email")
		return nil
	}

	return c
}
