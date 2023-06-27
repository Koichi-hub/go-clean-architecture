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

func (taskRepo *TaskRepo) Create(task entities.Task) (uint, error) {
	taskModel := fromTaskToTaskModel(task)

	err := taskRepo.db.Create(&taskModel).Error

	return taskModel.Id, err
}

func (taskRepo *TaskRepo) GetById(sessionId string, taskId uint) (entities.Task, error) {
	var taskModel models.TaskModel
	err := taskRepo.db.First(&taskModel, "session_id = ? AND id = ?", sessionId, taskId).Error
	if err != nil {
		return entities.Task{}, err
	}

	task := fromTaskModelToTask(taskModel)

	return task, nil
}

func (taskRepo *TaskRepo) GetAll(sessionId string) ([]entities.Task, error) {
	var tasksModels []models.TaskModel
	err := taskRepo.db.Find(&tasksModels, "session_id = ?", sessionId).Error
	if err != nil {
		return []entities.Task{}, err
	}

	tasks := make([]entities.Task, len(tasksModels))
	for i, taskModel := range tasksModels {
		tasks[i] = fromTaskModelToTask(taskModel)
	}

	return tasks, nil
}

func (taskRepo *TaskRepo) Update(task entities.Task) error {
	taskModel := fromTaskToTaskModel(task)
	err := taskRepo.db.Save(&taskModel).Error

	return err
}

func (taskRepo *TaskRepo) Delete(sessionId string, taskId uint) error {
	err := taskRepo.db.Delete(&models.TaskModel{}, "session_id = ? AND id = ?", sessionId, taskId).Error

	return err
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

func fromTaskModelToTask(taskModel models.TaskModel) entities.Task {
	return entities.Task{
		SessionId:   taskModel.SessionId,
		Id:          taskModel.Id,
		Title:       taskModel.Title,
		Description: taskModel.Description,
		Completed:   taskModel.Completed,
		CreatedAt:   taskModel.CreatedAt,
		UpdatedAt:   taskModel.UpdatedAt,
	}
}
