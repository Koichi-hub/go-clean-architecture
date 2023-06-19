package controllers

import (
	"go-clean-architecture/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase interfaces.TaskUseCase
}

func NewTaskController(taskUseCase interfaces.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

func (taskController *TaskController) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/tasks")
	g.POST("/", taskController.create)
	g.GET("/:taskId", taskController.getById)
	g.GET("/", taskController.getAll)
	g.PATCH("/complete", taskController.complete)
	g.PATCH("/", taskController.update)
	g.DELETE("/", taskController.delete)
}

func (taskController *TaskController) create(ctx *gin.Context) {

}

func (taskController *TaskController) getById(ctx *gin.Context) {

}

func (taskController *TaskController) getAll(ctx *gin.Context) {

}

func (taskController *TaskController) complete(ctx *gin.Context) {

}

func (taskController *TaskController) update(ctx *gin.Context) {

}

func (taskController *TaskController) delete(ctx *gin.Context) {

}
