package main

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "54789b648b06d1"
	password = "e22c265a82a8f0"
)

func main() {
	from := "greyrat999@mail.ru"
	to := "americano999level@mail.ru"
	subject := "this is a test email"
	plaintext := "this is a bode of the email"
	html := `<h1>hello</h1><p>this is email</p>`
	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetHeader("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	// msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
