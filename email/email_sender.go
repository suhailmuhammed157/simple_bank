package email

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

type MailSender interface {
	SendEmail(from string, to string, subject string, secretCode string) error
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

func (emailSender *EmailSender) SendEmail(from string, to string, subject string, secretCode string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	// m.Attach("/home/Alex/lolcat.jpg")

	// Format the message with the verification code
	body := fmt.Sprintf(
		`Hello, Thank you for registering with us. Please verify your email with the below code </br> <b>%v</b>.`,
		secretCode,
	)

	// Set the email body
	m.SetBody("text/html", body)

	d := gomail.NewDialer(emailSender.emailHost, int(emailSender.port), emailSender.username, emailSender.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
