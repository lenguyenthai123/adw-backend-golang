package mail

import (
	"backend-golang/config/models"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

type Mailer struct {
	mailConfig *models.MailConfig
}

func NewMailer(mailConfig *models.MailConfig) *Mailer {
	return &Mailer{mailConfig: mailConfig}
}

type Config struct {
	MailFrom   string
	MailServer string
	MailPort   int
	MailPass   string
}

func (m *Mailer) Init() Config {
	return Config{
		MailFrom:   m.mailConfig.MailFrom,
		MailServer: m.mailConfig.MailServer,
		MailPort:   m.mailConfig.MailPort,
		MailPass:   m.mailConfig.MailPass,
	}
}

func (m *Mailer) SendMail(to, subject, templatePath string, data interface{}) error {
	cfg := m.Init()

	if len(to) == 0 || len(subject) == 0 || len(templatePath) == 0 {
		return errors.New("to, subject, templatePath can not empty")
	}

	body, err := parseTemplate(templatePath, data)
	if err != nil {
		return err
	}

	var messages []string

	messages = append(messages, "From: Trekkstay<"+cfg.MailFrom+">\r")
	messages = append(messages, "To: "+to+"\r")
	messages = append(messages, "Subject: "+subject+"\r")
	messages = append(messages, "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n")
	messages = append(messages, body+"\r")

	msg := []byte(strings.Join(messages, "\n"))
	mailAuth := fmt.Sprintf("%s:%d", cfg.MailServer, cfg.MailPort)

	err = smtp.SendMail(mailAuth,
		smtp.PlainAuth("", cfg.MailFrom, cfg.MailPass, cfg.MailServer), cfg.MailFrom, []string{to}, msg)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Helper function help you bind data to the template
func parseTemplate(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
