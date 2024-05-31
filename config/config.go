package config

import (
	"fmt"
	"log"

	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=user password=password dbname=tarefas_db port=5433 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connection established")
	migrate()
}

func migrate() {

	err := DB.AutoMigrate(&model.Tarefa{}, &model.SMTP{}, &model.Period{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	fmt.Println("Database migration completed")
}
