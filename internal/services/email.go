package services

import (
	"log"
	"net/smtp"

	"anilkhadka.com.np/task-management/conf"
)

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (e *EmailService) SendEmail(recipient string, body string) error {
	// Sender's email address
	from := conf.EnvConfigs.EmailSender
	password := conf.EnvConfigs.EmailPassword

	// Recipient's email address
	to := recipient

	// SMTP server and port
	smtpServer := conf.EnvConfigs.EmailServer
	smtpPort := conf.EnvConfigs.EmailServerSMTPPort

	// Message content
	subject := "Your Email Subject"
	message := "Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body

	// Create the authentication credentials
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Connect to the SMTP server
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	log.Println("Email sent successfully.")
	return nil
}
