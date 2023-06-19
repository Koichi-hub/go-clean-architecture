package interfaces

import (
	"go-clean-architecture/entities"
	"go-clean-architecture/usecases/dto"
)

type TaskUseCase interface {
	Create(dto.CreateTaskDto) (dto.TaskDto, error)
	GetById(sessionId string, taskId uint) (dto.TaskDto, error)
	GetAll(sessionId string) ([]dto.TaskDto, error)
	Complete(sessionId string, taskId uint) (dto.TaskDto, error)
	Update(dto.UpdateTaskDto) (dto.TaskDto, error)
	Delete(sessionId string, taskId uint) error
}

type TaskRepo interface {
	Save(entities.Task) (entities.Task, error)
	GetById(sessionId string, taskId uint) (entities.Task, error)
	GetAll(sessionId string) ([]entities.Task, error)
	Complete(sessionId string, taskId uint) (entities.Task, error)
	Delete(sessionId string, taskId uint) error
}
