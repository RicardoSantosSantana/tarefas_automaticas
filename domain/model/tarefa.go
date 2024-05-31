package model

import "gorm.io/gorm"

type Tarefa struct {
	gorm.Model
	URL     string
	Hour    int
	*Period `gorm:"embedded"`
}

type Period struct {
	Daily    bool
	Weekly   bool
	Annually bool
}

type SMTP struct {
	Port     int
	Url      string
	User     string
	Password string
}

func NewPeriod(Daily, Weekly, Annually bool) *Period {
	return &Period{Daily, Weekly, Annually}
}

func NewTarefa(URL string, Hour int, Period *Period) *Tarefa {
	return &Tarefa{URL: URL, Hour: Hour, Period: Period}
}

func NewSMTP(Port int, Url string, User string, Password string) *SMTP {
	return &SMTP{Port, Url, User, Password}
}
