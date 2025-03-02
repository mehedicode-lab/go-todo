package main

import (
	"go-todo/config"
	"go-todo/internal/domain/models"
	"go-todo/internal/domain/services"
	"go-todo/internal/infastructure/datatabase/rds"
	"go-todo/internal/repository"
	"go-todo/internal/transport/http/handlers"
	"go-todo/internal/transport/http/routes"
	"log"
)

func main() {
	loadCnf, _ := config.LoadConfig()
	db := rds.ConnectDB(loadCnf.Rds)
	db.AutoMigrate(&models.Todo{})

	todoRepo := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	r := routes.SetupRoutes(todoHandler)

	log.Println("Server is running on :" + loadCnf.Server.Port)
	r.Run(":" + loadCnf.Server.Port)
}
