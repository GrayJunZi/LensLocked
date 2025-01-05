package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "smtp-mail.outlook.com"
	port     = 587
	username = "grayjunzi@outlook.com"
	password = "vyrjktlqencfjheor"
)

func main() {
	from := "grayjunzi@outlook.com"
	to := "grayjunzi@163.com"
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

	dialer := mail.NewDialer(host, port, username, password)

	/*
		sender, err := dialer.Dial()
		if err != nil {
			panic(err)
		}
		defer sender.Close()
		sender.Send(from, []string{to}, msg)
	*/

	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
