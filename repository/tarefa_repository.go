package repository

import (
	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"gorm.io/gorm"
)

type TarefaRepository struct {
	db *gorm.DB
}

func NewTarefaRepository(db *gorm.DB) *TarefaRepository {
	return &TarefaRepository{db: db}
}

func (r *TarefaRepository) Save(t *model.Tarefa) bool {
	result := r.db.Create(t)
	return result.Error == nil
}
