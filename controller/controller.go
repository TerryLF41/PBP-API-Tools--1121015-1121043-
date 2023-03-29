package controller

import (
	"time"

	"github.com/go-co-op/gocron"
	"gopkg.in/gomail.v2"
)

// GoCron
func runScheduler() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Days().Do(func() {
		//Func Kirim Email
	})

	s.StartBlocking()
}

// GoMail
func sendMail() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "EMAIL ADMIN")
	msg.SetHeader("To", "EMAIL USER")
	msg.SetHeader("Subject", "TITLE BERITA")
	msg.SetBody("text/html", "<p>ISI BERITA</p>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "EMAIL ADMIN", "PASSWORD EMAIL ADMIN")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
