package controllers

import (
	"go-clean-architecture/controllers/api_dto"
	"go-clean-architecture/usecases/dto"
	"go-clean-architecture/usecases/interfaces"
	"net/http"
	"strconv"

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
	g.POST("", taskController.Create)
	g.GET(":taskId", taskController.GetById)
	g.GET("", taskController.GetAll)
	g.PATCH("complete", taskController.Complete)
	g.PUT("", taskController.Update)
	g.DELETE("", taskController.Delete)
}

func (taskController *TaskController) Create(ctx *gin.Context) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	var createTaskRequest api_dto.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&createTaskRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, newHttpError(err.Error()))
		return
	}

	taskDto := fromCreateTaskRequestToCreateTaskDto(createTaskRequest)
	taskDto.SessionId = sessionId

	taskId, err := taskController.taskUseCase.Create(taskDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, taskId)
}

func (taskController *TaskController) GetById(ctx *gin.Context) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	param_id := ctx.Param("id")
	taskId, err := strconv.Atoi(param_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, newHttpError(err.Error()))
		return
	}

	taskDto, err := taskController.taskUseCase.GetById(sessionId, uint(taskId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	taskResponse := fromTaskDtoToTaskResponse(taskDto)
	ctx.JSON(http.StatusOK, taskResponse)
}

func (taskController *TaskController) GetAll(ctx *gin.Context) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	tasksDto, err := taskController.taskUseCase.GetAll(sessionId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	tasksResponse := make([]api_dto.TaskResponse, len(tasksDto))
	for i, taskDto := range tasksDto {
		taskResponse := fromTaskDtoToTaskResponse(taskDto)
		tasksResponse[i] = taskResponse
	}

	ctx.JSON(http.StatusOK, tasksResponse)
}

func (taskController *TaskController) Complete(ctx *gin.Context) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	param_id := ctx.Param("id")
	taskId, err := strconv.Atoi(param_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, newHttpError(err.Error()))
		return
	}

	if err := taskController.taskUseCase.Complete(sessionId, uint(taskId)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (taskController *TaskController) Update(ctx *gin.Context) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	param_id := ctx.Param("id")
	taskId, err := strconv.Atoi(param_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, newHttpError(err.Error()))
		return
	}

	var updateTaskRequest api_dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&updateTaskRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, newHttpError(err.Error()))
		return
	}

	updateTaskDto := fromUpdateTaskRequestToUpdateTaskDto(updateTaskRequest)
	updateTaskDto.SessionId = sessionId
	updateTaskDto.Id = uint(taskId)

	if err := taskController.taskUseCase.Update(updateTaskDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (taskController *TaskController) Delete(ctx *gin.Context) {
	sessionId, err := getSessionId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	param_id := ctx.Param("id")
	taskId, err := strconv.Atoi(param_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, newHttpError(err.Error()))
		return
	}

	err = taskController.taskUseCase.Delete(sessionId, uint(taskId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, newHttpError(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func fromCreateTaskRequestToCreateTaskDto(createTaskRequest api_dto.CreateTaskRequest) dto.CreateTaskDto {
	return dto.CreateTaskDto{
		Title:       createTaskRequest.Title,
		Description: createTaskRequest.Description,
	}
}

func fromUpdateTaskRequestToUpdateTaskDto(updateTaskRequest api_dto.UpdateTaskRequest) dto.UpdateTaskDto {
	return dto.UpdateTaskDto{
		Title:       updateTaskRequest.Title,
		Description: updateTaskRequest.Description,
	}
}

func fromTaskDtoToTaskResponse(taskDto dto.TaskDto) api_dto.TaskResponse {
	return api_dto.TaskResponse{
		SessionId:   taskDto.SessionId,
		Id:          taskDto.Id,
		Title:       taskDto.Title,
		Description: taskDto.Description,
		Completed:   taskDto.Completed,
		CreatedAt:   taskDto.CreatedAt,
		UpdatedAt:   taskDto.UpdatedAt,
	}
}
