package mail

import (
	"net/smtp"
	"strconv"
)

type Mailer interface {
	SendEmail(to []string, subject string, body string) error
	VerifyEmailTemplate(email string, link string) string
}

type MailService struct {
	Host     string
	Port     int
	name     string
	email    string
	Password string
}

func NewMailSender(name string, fromEmailAddress string, fromEmailPassword string) Mailer {
	return &MailService{
		Host:     "smtp.gmail.com",
		Port:     587,
		name:     name,
		email:    fromEmailAddress,
		Password: fromEmailPassword,
	}
}

func (m *MailService) SendEmail(to []string, subject string, body string) error {
	receiver := to
	// Sender data.
	from := m.email
	password := m.Password

	// smtp server configuration.
	smtpHost := m.Host
	smtpPort := m.Port

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte("To: " + receiver[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), // Convert Int to String
		auth,
		from,
		receiver,
		msg)

	if err != nil {
		return err
	}

	return nil
}

func (m *MailService) VerifyEmailTemplate(email string, link string) string {
	return `
		<p>Hi ` + email + `</p>
		<p>Please verify your email by using the following link: <a href="` + link + `">Verify Email</a></p>
		<p>Regards,</p>
		<p>Instagram Team</p>
	`
}
