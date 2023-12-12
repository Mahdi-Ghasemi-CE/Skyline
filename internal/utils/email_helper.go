package utils

import (
	"github.com/go-mail/mail"
)

const (
	smtpAuthAddress = "smtp.gmail.com"
	smtpPortAddress = 587
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

func SendEmail(
	subject string,
	content string,
	to string,
	attachFiles []string,
) error {
	m := mail.NewMessage()

	m.SetHeader("From", AppConfig.EmailSenderAddress)

	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", content)

	d := mail.NewDialer(smtpAuthAddress, smtpPortAddress, AppConfig.EmailSenderAddress, AppConfig.EmailSenderPassword)

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {

		return err
	}
	return nil
}
