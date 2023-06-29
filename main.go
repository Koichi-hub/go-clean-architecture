package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/controllers"
	"go-clean-architecture/databases/mysql/db"
	"go-clean-architecture/databases/mysql/repos"
	"go-clean-architecture/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	cfg := config.LoadConfig(".env")

	// database
	db := db.Connect(cfg)

	// repositories
	taskRepo := repos.NewTaskRepo(db)

	// usecases
	taskUseCase := usecases.NewTaskUseCase(taskRepo)

	// controllers
	taskController := controllers.NewTaskController(taskUseCase)

	// http server
	r := gin.Default()

	taskController.RegisterRoutes(r)

	address := fmt.Sprintf("%s:%d", cfg.HOST, cfg.PORT)
	r.Run(address)
}
