package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	DefaultSender = "support@lenslocked.com"
)

type Email struct {
	From      string
	To        string
	Subject   string
	PlainText string
	HTML      string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type EmailService struct {
	DefaultSender string

	dialer *mail.Dialer
}

func NewEmailService(config SMTPConfig) *EmailService {
	return &EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
}

func (e *EmailService) Send(email Email) error {
	msg := mail.NewMessage()

	e.setFrom(msg, email)

	msg.SetHeader("From", email.From)
	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)

	if email.PlainText != "" {
		msg.SetBody("text/plain", email.PlainText)
	}

	if email.HTML != "" {
		msg.AddAlternative("text/html", email.HTML)
	}

	if err := e.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("send: %s", err)
	}

	return nil
}

func (e *EmailService) ForgotPassword(to, resetURL string) error {
	email := Email{
		Subject:   "Reset your password",
		To:        to,
		PlainText: "To reset your password, please visit the following link:" + resetURL,
		HTML:      `<p>To reset your password, pleanse visit the following link: <a href="` + resetURL + `">` + resetURL + `</a></p>`,
	}

	err := e.Send(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}

	return nil
}

func (e *EmailService) setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	case e.DefaultSender != "":
		from = e.DefaultSender
	default:
		from = DefaultSender
	}
	msg.SetHeader("From", from)
}
