package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending mail:", err)
		return err
	}
	fmt.Println("Mail sent successfully to", to)
	return nil
}
