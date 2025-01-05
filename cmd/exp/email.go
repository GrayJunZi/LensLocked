package main

import (
	"os"

	"github.com/go-mail/mail/v2"
)

func main() {
	from := "admin@email.com"
	to := "test@email.com"
	subject := "This is a test email"

	plaintext := "This is the body of the email"
	html := `<h1>Hello, World!</h1>`

	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)
}
