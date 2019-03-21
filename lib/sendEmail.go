package lib

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
)

type Mail struct {
	from    string   // veloss <verification@gmail.com>
	to      []string // email
	subject string   // veloss ${emailKeywords.text}
	body    string   //html
}

func NewRequest(to []string, subject, body string) *Mail {
	return &Mail{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (m *Mail) SendEmail() bool {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASS"), "smtp.gmail.com")
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + m.subject + "!\n"
	msg := []byte(subject + mime + "\n" + m.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, m.from, m.to, msg); err != nil {
		return false
	}
	return true
}

func (m *Mail) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	m.body = buf.String()
	return nil
}
