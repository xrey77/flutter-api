package middlewares

import (
	"crypto/tls"

	"gopkg.in/mail.v2"
)

func SendMail(msg string, sub string, xmail string) int64 {
	m := mail.NewMessage()
	m.SetHeader("From", "rey107@gmail.com")
	m.SetHeader("To", xmail)
	m.SetHeader("Subject", sub)
	m.SetBody("text/html", msg)
	d := mail.NewDialer("smtp.gmail.com", 587, "rey107@gmail.com", "Reynald@88.88")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return 0
	}
	return 1
}
