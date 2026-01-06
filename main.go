package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Livro struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"column:title; not null"`
	Author string `json:"author" gorm:"column:author; not null"`
	PublishedYear int `json:"published_year" gorm:"column:published_year; not null"`
}

func main() {
	host:= os.Getenv("DB_HOST")
	user:= os.Getenv("DB_USER")
	password:= os.Getenv("DB_PASSWORD")
	dbName:= os.Getenv("DB_NAME")
	dbPort:= os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
    host, user, password, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados:", err)
	}

	db.AutoMigrate(&Livro{})
	fmt.Println("Migração concluída!")

	fmt.Println("Aguardando o banco de dados subir...")
    time.Sleep(5 * time.Second)
}