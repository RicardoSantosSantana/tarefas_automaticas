package service

import (
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/interfaces"
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
)

type TarefaService struct {
	repo interfaces.Tarefa
	smtp interfaces.SMTP
}

func NewTarefaService(repo interfaces.Tarefa, smtp interfaces.SMTP) *TarefaService {
	return &TarefaService{repo: repo, smtp: smtp}
}

func (s *TarefaService) SalvarAgendamento(message string, email string, URL string, Hour int, Daily bool, Weekly bool, Annually bool) bool {
	periodo := model.NewPeriod(Daily, Weekly, Annually)
	tarefa := model.NewTarefa(URL, Hour, periodo)
	if s.repo.Save(tarefa) {
		return s.smtp.Send(email, message)
	}
	return false
}
