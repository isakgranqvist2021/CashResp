package utils

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

type Mail struct {
	Receivers []string
	Message   string
}

func SendMail(m *Mail) {
	// read .env file in root directory
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	from := os.Getenv("GMAIL_EMAIL")
	password := os.Getenv("GMAIL_PASSWORD")

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", m.Receivers...)
	mail.SetHeader("Subject", "Cashresp Verification Code")
	mail.SetBody("text/html", m.Message)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
