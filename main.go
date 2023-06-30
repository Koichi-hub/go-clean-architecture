package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/controllers"
	"go-clean-architecture/controllers/middlewares"
	"go-clean-architecture/databases/mysql/db"
	"go-clean-architecture/databases/mysql/repos"
	"go-clean-architecture/usecases"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	cfg := config.LoadConfig()

	// database
	db := db.Connect(cfg)

	// repositories
	taskRepo := repos.NewTaskRepo(db)

	// usecases
	taskUseCase := usecases.NewTaskUseCase(taskRepo)

	// controllers
	sessionController := controllers.NewSessionController()
	taskController := controllers.NewTaskController(taskUseCase)

	// middlewares
	sessionMiddleware := middlewares.NewSessionMiddleware()

	// http server
	if cfg.MODE == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	sessionController.RegisterRoutes(r)
	taskController.RegisterRoutes(r, sessionMiddleware)

	log.Println("Starting http-server...")
	address := fmt.Sprintf("%s:%d", cfg.HOST, cfg.PORT)
	r.Run(address)
}
