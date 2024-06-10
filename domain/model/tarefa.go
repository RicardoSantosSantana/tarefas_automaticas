package model

import "gorm.io/gorm"

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

func NewEMail(From, To, Subject, Body string) *Email {
	return &Email{From, To, Subject, Body}
}

func NewTarefa(URL string, Hour int, Daily, Weekly, Annually bool) *Tarefa {
	return &Tarefa{URL: URL, Hour: Hour, Daily: Daily, Weekly: Weekly, Annually: Annually}
}

func NewSMTP(Port string, Url string, User string, Password string) *SMTP {
	return &SMTP{Port, Url, User, Password}
}

func (r *Email) MessageBody() string {

	from := "From: " + r.From + "\r\n"
	to := "To: " + r.To + "\r\n"
	subject := "Subject: " + r.Subject + "\r\n"
	body := r.Body

	return from + to + subject + body

}
