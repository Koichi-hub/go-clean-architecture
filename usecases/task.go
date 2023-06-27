package usecases

import (
	"go-clean-architecture/entities"
	"go-clean-architecture/usecases/dto"
	"go-clean-architecture/usecases/interfaces"
)

type TaskUseCase struct {
	taskRepo interfaces.TaskRepo
}

func NewTaskUseCase(taskRepo interfaces.TaskRepo) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (taskUseCase *TaskUseCase) Create(createTaskDto dto.CreateTaskDto) (uint, error) {
	task := fromCreateTaskDtoToTask(createTaskDto)

	return taskUseCase.taskRepo.Create(task)
}

func (taskUseCase *TaskUseCase) GetById(sessionId string, taskId uint) (dto.TaskDto, error) {
	task, err := taskUseCase.taskRepo.GetById(sessionId, taskId)
	if err != nil {
		return dto.TaskDto{}, err
	}

	taskDto := fromTaskToTaskDto(task)
	return taskDto, nil
}

func (taskUseCase *TaskUseCase) GetAll(sessionId string) ([]dto.TaskDto, error) {
	tasks, err := taskUseCase.taskRepo.GetAll(sessionId)
	if err != nil {
		return []dto.TaskDto{}, err
	}

	tasksDto := make([]dto.TaskDto, len(tasks))
	for i, task := range tasks {
		taskDto := fromTaskToTaskDto(task)
		tasksDto[i] = taskDto
	}
	return tasksDto, nil
}

func (taskUseCase *TaskUseCase) Complete(sessionId string, taskId uint) error {
	task, err := taskUseCase.taskRepo.GetById(sessionId, taskId)
	if err != nil {
		return err
	}

	task.Completed = true

	return taskUseCase.taskRepo.Update(task)
}

func (taskUseCase *TaskUseCase) Update(updateTaskDto dto.UpdateTaskDto) error {
	task := fromUpdateTaskDtoToTask(updateTaskDto)

	return taskUseCase.taskRepo.Update(task)
}

func (taskUseCase *TaskUseCase) Delete(sessionId string, taskId uint) error {
	return taskUseCase.taskRepo.Delete(sessionId, taskId)
}

func fromCreateTaskDtoToTask(taskDto dto.CreateTaskDto) entities.Task {
	return entities.Task{
		SessionId:   taskDto.SessionId,
		Title:       taskDto.Title,
		Description: taskDto.Description,
	}
}

func fromUpdateTaskDtoToTask(taskDto dto.UpdateTaskDto) entities.Task {
	return entities.Task{
		SessionId:   taskDto.SessionId,
		Id:          taskDto.Id,
		Title:       taskDto.Title,
		Description: taskDto.Description,
	}
}

func fromTaskToTaskDto(task entities.Task) dto.TaskDto {
	return dto.TaskDto{
		SessionId:   task.SessionId,
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
