package main

import (
	"biblioteca-api/database"
	"biblioteca-api/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


func main() {

	database.ConectarComBancoDeDados()

	database.DB.AutoMigrate(&models.Livro{})
	
	router := gin.Default()	

	fmt.Println("Aguardando o banco de dados subir...")
    time.Sleep(5 * time.Second)

	router.Run(":8080")
}