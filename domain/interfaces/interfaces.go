package interfaces

import "github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"

type Tarefa interface {
	Save(*model.Tarefa) bool
}

type SMTP interface {
	Save() bool
	Get() model.SMTP
	Send() bool
}
