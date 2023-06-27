package repos

import (
	"go-clean-architecture/databases/mysql/models"
	"go-clean-architecture/entities"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (taskRepo *TaskRepo) Create(task entities.Task) error {
	taskModel := fromTaskToTaskModel(task)

	return taskRepo.db.Create(&taskModel).Error
}

func (taskRepo *TaskRepo) GetById(sessionId string, taskId uint) (entities.Task, error) {
	return entities.Task{}, nil
}

func (taskRepo *TaskRepo) GetAll(sessionId string) ([]entities.Task, error) {
	return []entities.Task{}, nil
}

func (taskRepo *TaskRepo) Update(entities.Task) error {
	return nil
}

func (taskRepo *TaskRepo) Delete(sessionId string, taskId uint) error {
	return nil
}

func fromTaskToTaskModel(task entities.Task) models.TaskModel {
	return models.TaskModel{
		SessionId:   task.SessionId,
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
