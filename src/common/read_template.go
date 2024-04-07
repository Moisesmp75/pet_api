package common

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"pet_api/src/models"
	"text/template"
	"time"
)

type EmailData struct {
	RecipientName string
	DateTime      string
	NewPassword   string
}

func SendResetPasswordEmail(user models.User, newPassword string) error {

	emailData := EmailData{
		RecipientName: user.Name + " " + user.LastName,
		DateTime:      time.Now().Format("02-01-2006 15:04:05"),
		NewPassword:   newPassword,
	}

	tmpl, err := template.ParseFiles("src/common/templates/reset_password.html")
	if err != nil {
		return fmt.Errorf("error loading HTML template: %w", err)
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, emailData); err != nil {
		return fmt.Errorf("error rendering HTML template: %w", err)
	}

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Recover Password"
	msg := []byte(fmt.Sprintf("Subject: %s\n", subject) +
		"Content-Type: text/html; charset=UTF-8\n" +
		"\n" + tpl.String())

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{user.Email}, msg)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
