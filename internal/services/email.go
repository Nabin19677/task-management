package services

import (
	"log"
	"net/smtp"
	"strconv"

	"anilkhadka.com.np/task-management/conf"
)

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (e *EmailService) SendEmail(recipient string, body string) error {
	from := conf.EnvConfigs.Email
	password := conf.EnvConfigs.EmailPassword

	// Recipient's email address
	to := recipient

	// SMTP server and port
	smtpServer := "smtp.mail.yahoo.com"
	smtpPort := 587

	// Message content
	subject := "Your Email Subject"
	message := "Subject: " + subject + "\r\n" +
		"\r\n" +
		body

	// Create the authentication credentials
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Connect to the SMTP server
	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	log.Println("Email sent successfully.")
	return nil
}
