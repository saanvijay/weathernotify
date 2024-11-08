package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendweatherNotifyMail(subject string, body string) {

	to := []string{os.Getenv("TO_EMAIL_ID")}
	auth := smtp.PlainAuth(
		"",
		os.Getenv("FROM_EMAIL_ID"),
		os.Getenv("FROM_EMAIL_APP_PASS"),
		"smtp.gmail.com",
	)
	msg := "Subject: " + subject + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("FROM_EMAIL_ID"),
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err)
		return
	}
}
