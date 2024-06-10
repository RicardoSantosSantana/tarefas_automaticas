package main

import (
	"github.com/RicardoSantosSantana/tarefas_automaticas/config"
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"github.com/RicardoSantosSantana/tarefas_automaticas/repository"
)

func main() {
	// SMTP
	port := "587"
	url := "smtp.needsolution.com.br"
	user := "ricardo@needsolution.com.br"
	password := "senha001"
	SMTP := model.NewSMTP(port, url, user, password)

	// EMAIL
	from := "ricardo@needsolution.com.br"
	to := "rssantan@gmail.com"
	subject := "Mensagem de Teste 2"
	message := "Esta é uma mensagem de teste com envio via Golang 2"
	EMAIL := model.NewEMail(from, to, subject, message)

	smtpRepo := repository.NewSMTPRepository(config.DB, SMTP, EMAIL)
	smtpRepo.Send()

}

func Main2() {
	//config.InitDB()

	//tarefaRepo := repository.NewTarefaRepository(config.DB)

	// tarefaService := service.NewTarefaService(tarefaRepo, smtpRepo)

	// result := tarefaService.SalvarAgendamento(SMTP, EMAIL, "http://example.com", 10, true, false, false)

	// fmt.Printf("SalvarAgendamento result: %v\n", result)
}
