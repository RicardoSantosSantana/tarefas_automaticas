package config

import (
	"fmt"
	"log"
	"os"

	"github.com/RicardoSantosSantana/tarefas_automaticas/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDBURL() string {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	charset := os.Getenv("CHARSET")

	// Montar a string de conexão
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, port, database, charset)

	return dbURL
}

func InitDB() {
	//dsn := "casaos:casaos@tcp(192.168.0.4:3306)/casaos?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	dsn := CreateDBURL()

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connection established\n", dsn)
	migrate()
}

func migrate() {
	err := DB.AutoMigrate(&model.Tarefa{}, &model.SMTP{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	fmt.Println("Database migration completed")
}
