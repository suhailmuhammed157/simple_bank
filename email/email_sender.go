package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type MailSender interface {
	SendEmail(from string, to string) error
}

type EmailSender struct {
	emailHost string
	port      int32
	username  string
	password  string
}

func NewEmailSender(emailHost string, port int32, username string, password string) MailSender {

	return &EmailSender{
		emailHost: emailHost,
		port:      port,
		username:  username,
		password:  password,
	}
}

func (emailSender *EmailSender) SendEmail(from string, to string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(emailSender.emailHost, int(emailSender.port), emailSender.username, emailSender.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
