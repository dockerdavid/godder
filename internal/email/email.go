package email

import (
	"crypto/tls"
	"godder/internal/config"
	"log"

	gomail "gopkg.in/mail.v2"
)

func SendMail(body string) {
	m := gomail.NewMessage()

	m.SetHeader("From", config.Config.Godder.Email.From)

	m.SetHeader("To", config.Config.Godder.Email.To)

	m.SetHeader("Subject", "Godder Alert")

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(config.Config.Godder.Email.Host, config.Config.Godder.Email.Port, config.Config.Godder.Email.From, config.Config.Godder.Email.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
}
