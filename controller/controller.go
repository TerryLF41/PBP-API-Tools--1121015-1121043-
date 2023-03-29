package controller

import (
	"time"

	"github.com/go-co-op/gocron"
	"gopkg.in/gomail.v2"
)

// GoCron
func runScheduler() {
	schedule := gocron.NewScheduler(time.UTC)

	schedule.Every(1).Days().Do(func() {
		//Func Kirim Email
	})

	schedule.StartBlocking()
}

// GoMail
func sendMail(user User, news Berita) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "EMAIL ADMIN")
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", news.Title)
	msg.SetBody("text/html", "<p>"+news.Isi+"</p>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "EMAIL ADMIN", "PASSWORD EMAIL ADMIN")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func getTodayNews() {
	db := Connect()
	defer db.Close()
}
