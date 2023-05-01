package service

import (
	"encoding/base64"
	"net/smtp"
	"strings"
)

type EmailSender struct {
	c *smtp.Client
}

func NewEmailSender(c *smtp.Client) *EmailSender {
	return &EmailSender{c: c}
}

func (m EmailSender) SendEmail(subj, body, from string, to ...string) error {
	if err := m.c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err := m.c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := m.c.Data()

	go func() {
		defer w.Close()
		w.Write(
			[]byte("To: " + strings.Join(to, ",") + "\r\n" +
				"Subject: " + subj + "\r\n" +
				"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
				"Content-Transfer-Encoding: base64\r\n" +
				"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))),
		)
	}()
	return err
}
