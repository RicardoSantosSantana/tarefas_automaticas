package main

import (
	"fmt"

	"github.com/RicardoSantosSantana/tarefas_automaticas/config"
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/service"
	"github.com/RicardoSantosSantana/tarefas_automaticas/repository"
)

func main() {
	config.InitDB()

	tarefaRepo := repository.NewTarefaRepository(config.DB)
	smtpRepo := repository.NewSMTPRepository(config.DB)

	tarefaService := service.NewTarefaService(tarefaRepo, smtpRepo)

	result := tarefaService.SalvarAgendamento("This is a test message", "test@example.com", "http://example.com", 10, true, false, false)
	fmt.Printf("SalvarAgendamento result: %v\n", result)
}
