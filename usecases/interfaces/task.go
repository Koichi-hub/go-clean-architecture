package interfaces

import (
	"go-clean-architecture/entities"
	"go-clean-architecture/usecases/dto"
)

type TaskUseCase interface {
	Create(dto.CreateTaskDto) error
	GetById(sessionId string, taskId uint) (dto.TaskDto, error)
	GetAll(sessionId string) ([]dto.TaskDto, error)
	Complete(sessionId string, taskId uint) error
	Update(dto.UpdateTaskDto) error
	Delete(sessionId string, taskId uint) error
}

type TaskRepo interface {
	Create(entities.Task) error
	GetById(sessionId string, taskId uint) (entities.Task, error)
	GetAll(sessionId string) ([]entities.Task, error)
	Update(entities.Task) error
	Delete(sessionId string, taskId uint) error
}
