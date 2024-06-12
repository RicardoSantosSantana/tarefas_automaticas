package repository

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"gorm.io/gorm"
)

type SMTPRepository struct {
	db    *gorm.DB
	SMTP  *model.SMTP
	EMAIL *model.Email
}

func NewSMTPRepository(db *gorm.DB, smtp *model.SMTP, email *model.Email) *SMTPRepository {
	return &SMTPRepository{db: db, SMTP: smtp, EMAIL: email}
}

func (r *SMTPRepository) Save() bool {
	result := r.db.Create(&r.SMTP)
	return result.Error == nil
}

func (r *SMTPRepository) Get() model.SMTP {
	var smtp model.SMTP
	r.db.First(&smtp)
	return smtp
}

func (r *SMTPRepository) Send() bool {
	// Configurar a autenticação SMTP
	auth := smtp.PlainAuth("", r.SMTP.User, r.SMTP.Password, r.SMTP.ServerURL)

	// Compor o corpo do email
	emailBody := r.EMAIL.MessageBody()

	// Conectar ao servidor SMTP
	conn, err := smtp.Dial(r.SMTP.ServerURL + ":" + r.SMTP.Port)
	if err != nil {
		fmt.Println("Falha ao conectar ao servidor SMTP:", err)
		return false
	}
	defer conn.Close()

	// Configurar a conexão TLS usando STARTTLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Use com cuidado em produção
		ServerName:         r.SMTP.ServerURL,
	}

	if err = conn.StartTLS(tlsConfig); err != nil {
		fmt.Println("Falha ao iniciar a conexão TLS:", err)
		return false
	}

	// Autenticar com o servidor SMTP
	if err := conn.Auth(auth); err != nil {
		fmt.Println("Falha na autenticação SMTP:", err)
		return false
	}

	// Enviar o email
	if err := conn.Mail(r.SMTP.User); err != nil {
		fmt.Println("Falha ao definir o remetente:", err)
		return false
	}
	if err := conn.Rcpt(r.EMAIL.To); err != nil {
		fmt.Println("Falha ao definir o destinatário:", err)
		return false
	}

	// Abrir o corpo do email
	wc, err := conn.Data()
	if err != nil {
		fmt.Println("Falha ao abrir o corpo do email:", err)
		return false
	}
	defer wc.Close()

	// Escrever o corpo do email
	_, err = wc.Write([]byte(emailBody))
	if err != nil {
		fmt.Println("Falha ao escrever o corpo do email:", err)
		return false
	}

	fmt.Println("Email enviado com sucesso")
	return true
}
