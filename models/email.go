package models

import "github.com/go-mail/mail/v2"

const (
	DefaultSender = "support@lenslocked.com"
)

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
