package service

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailSender struct {
	c *smtp.Client
}

var r = strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

func NewEmailSender(c *smtp.Client) *EmailSender {
	return &EmailSender{c: c}
}

func (m EmailSender) SendEmail(subj, body, from string, to ...string) error {
	if m.c == nil {
		return fmt.Errorf("email client is not set. Try to define the environment variables")
	}

	if err := m.c.Mail(r.Replace(from)); err != nil {
		return err
	}

	for i := range to {
		to[i] = r.Replace(to[i])
		if err := m.c.Rcpt(to[i]); err != nil {
			return err
		}
	}

	w, err := m.c.Data()
	if err != nil {
		return err
	}

	go func() {
		defer w.Close()
		w.Write([]byte(
			"To: " + strings.Join(to, ",") + "\r\n" +
				"Subject: " + subj + "\r\n" +
				"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
				"Content-Transfer-Encoding: base64\r\n" +
				"\r\n" + base64.StdEncoding.EncodeToString([]byte(body)),
		))
	}()
	return nil
}
