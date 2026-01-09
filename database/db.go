package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é uma variável global que guardará a conexão
var DB *gorm.DB

func ConectarComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	fmt.Println("Aguardando o banco de dados subir...")
	time.Sleep(5 * time.Second)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados", err)
	}
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}