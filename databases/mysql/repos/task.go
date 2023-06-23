package repos

import "go-clean-architecture/entities"

type TaskRepo struct{}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (taskRepo *TaskRepo) Create(entities.Task) error {
	return nil
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
