package email

import (
	"crypto/tls"
	"fmt"
	"godder/internal/config"

	gomail "gopkg.in/mail.v2"
)

func SendMail(body string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Config.Godder.Email.From)

	// Set E-Mail receivers
	m.SetHeader("To", config.Config.Godder.Email.To)

	// Set E-Mail subject
	m.SetHeader("Subject", "Godder Alert")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer(config.Config.Godder.Email.Host, config.Config.Godder.Email.Port, config.Config.Godder.Email.From, config.Config.Godder.Email.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
