package main

import (
	"fmt"
	"os"

	"github.com/RicardoSantosSantana/tarefas_automaticas/config"
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/service"
	"github.com/RicardoSantosSantana/tarefas_automaticas/repository"
)

func main() {

	config.InitDB()

	EMAIL, err := model.NewEMail()
	if err != nil {
		for _, erros := range err {
			fmt.Println(erros.Error())
		}
		os.Exit(2)
	}

	SMTP, errSMTP := model.NewSMTP()

	if errSMTP != nil {
		for _, erros_smtp := range errSMTP {
			fmt.Println(erros_smtp.Error())
		}
		os.Exit(2)
	}

	smtpRepo := repository.NewSMTPRepository(config.DB, SMTP, EMAIL)

	tarefaRepo := repository.NewTarefaRepository(config.DB)

	tarefaService := service.NewTarefaService(tarefaRepo, smtpRepo)

	result := tarefaService.SalvarAgendamento("http://example.com", 10, true, false, false)

	fmt.Printf("SalvarAgendamento result: %v\n", result)

}
