package model

import (
	"errors"
	"os"
	"strings"

	"gorm.io/gorm"
)

type Tarefa struct {
	gorm.Model
	URL      string
	Hour     int
	Daily    bool
	Weekly   bool
	Annually bool
	Send     bool
	Error    string
}

type SMTP struct {
	Port      string
	ServerURL string
	User      string
	Password  string
}

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewEMail() (*Email, []error) {

	var errorList []error

	From := os.Getenv("EMAIL_FROM")
	To := os.Getenv("EMAIL_TO")
	Subject := os.Getenv("EMAIL_SUBJECT")
	Body := os.Getenv("EMAIL_MESSAGE")

	if strings.TrimSpace(From) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'EMAIL_FROM' não informada"))
	}

	if strings.TrimSpace(To) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'EMAIL_TO' não informada"))
	}

	if strings.TrimSpace(Subject) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'EMAIL_SUBJECT' não informada"))
	}

	if strings.TrimSpace(Body) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'EMAIL_BODY' não informada"))
	}

	if len(errorList) == 0 {
		return &Email{From, To, Subject, Body}, nil
	}

	return &Email{From, To, Subject, Body}, errorList
}

func NewTarefa(URL string, Hour int, Daily, Weekly, Annually bool) *Tarefa {
	return &Tarefa{URL: URL, Hour: Hour, Daily: Daily, Weekly: Weekly, Annually: Annually}
}

func NewSMTP() (*SMTP, []error) {

	var errorList []error

	Port := os.Getenv("SMTP_PORT")
	Url := os.Getenv("SMTP_URL")
	User := os.Getenv("SMTP_USER")
	Password := os.Getenv("SMTP_PASSWORD")

	if strings.TrimSpace(Port) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'SMTP_' não informada"))
	}

	if strings.TrimSpace(Url) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'SMTP_URL' não informada"))
	}

	if strings.TrimSpace(User) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'SMTP_USER' não informada"))
	}

	if strings.TrimSpace(Password) == "" {
		errorList = append(errorList, errors.New("variável de ambiente 'SMTP_PASSWORD' não informada"))
	}

	if len(errorList) == 0 {
		return &SMTP{Port: Port, ServerURL: Url, User: User, Password: Password}, nil
	}

	return &SMTP{
		Port:      Port,
		ServerURL: Url,
		User:      User,
		Password:  Password,
	}, errorList
}

func (r *Email) MessageBody() string {
	from := "From: " + r.From + "\r\n"
	to := "To: " + r.To + "\r\n"
	subject := "Subject: " + r.Subject + "\r\n"
	body := r.Body
	return from + to + subject + body
}
