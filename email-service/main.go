package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {
	// Configuration
	from := "yousafbhaikhan10@gmail.com"
	password := os.Getenv("pass")
	to := []string{"yousafbhaikhan10@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("Hello! I'm trying out smtp to send emails to recipients.")
	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)
	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
