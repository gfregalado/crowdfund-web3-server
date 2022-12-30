package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gfregalado/crowdfund-api/config"
	"github.com/gfregalado/crowdfund-api/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser
func SendEmail(user *models.DBResponse, data *EmailData, temp *template.Template, templateName string) error {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("could not load config", err)
	}

	// Sender data.

	/* from := config.EmailFrom
	smtpPass := config.SMTPPass
	smtpUser := config.SMTPUser
	to := user.Email
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	*/
	var body bytes.Buffer

	if err := temp.ExecuteTemplate(&body, templateName, &data); err != nil {
		log.Fatal("Could not execute template", err)
	}

	from := mail.NewEmail("info", config.EmailFrom)
	to := mail.NewEmail(user.Name, user.Email)
	subject := data.Subject
	plainTextContent := body.String()
	htmlContent := body.String()

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendGridApiKey)

	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
	}

	return nil
}

// ? Email template parser
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	fmt.Println("I am parsing templates...")

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
