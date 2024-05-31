package repository

import (
	"fmt"

	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"gorm.io/gorm"
)

type SMTPRepository struct {
	db *gorm.DB
}

func NewSMTPRepository(db *gorm.DB) *SMTPRepository {
	return &SMTPRepository{db: db}
}

func (r *SMTPRepository) Save(smtp model.SMTP) bool {
	result := r.db.Create(smtp)
	return result.Error == nil
}

func (r *SMTPRepository) Get() model.SMTP {
	var smtp model.SMTP
	r.db.First(&smtp)
	return smtp
}

func (r *SMTPRepository) Send(email string, message string) bool {
	// Implementação do envio de email fictício
	fmt.Printf("Sending email to %s with message: %s\n", email, message)
	return true
}
