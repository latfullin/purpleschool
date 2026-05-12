package mailing

import (
	"crypto/tls"
	"kilkenny/purpleschool/configs"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

type Sender struct {
	Name    string
	To      string
	Subject string
	HTML    string
}

func Send(r *Sender) error {
	config := configs.LoadConfig()

	e := &email.Email{
		To:      []string{r.To},
		From:    "Auth <Kilkenny94@yandex.ru>",
		Subject: r.Subject,
		HTML:    []byte(r.HTML),
		Headers: textproto.MIMEHeader{},
	}

	return e.SendWithTLS(
		config.SMTP.HOST+config.SMTP.PORT,
		smtp.PlainAuth(
			"",
			config.SMTP.FROM,
			config.SMTP.Password,
			config.SMTP.HOST,
		),
		&tls.Config{
			ServerName: config.SMTP.HOST,
		},
	)
}
