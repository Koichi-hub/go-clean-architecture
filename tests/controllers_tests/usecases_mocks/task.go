package usecases_mocks

import (
	"go-clean-architecture/usecases/dto"

	"github.com/stretchr/testify/mock"
)

type TaskUseCaseMock struct {
	mock.Mock
}

func NewTaskUseCaseMock() *TaskUseCaseMock {
	return &TaskUseCaseMock{}
}

func (taskUseCaseMock *TaskUseCaseMock) Create(createTaskDto dto.CreateTaskDto) (dto.TaskDto, error) {
	args := taskUseCaseMock.Called(createTaskDto)

	return args.Get(0).(dto.TaskDto), args.Error(1)
}

func (taskUseCaseMock *TaskUseCaseMock) GetById(sessionId string, taskId uint) (dto.TaskDto, error) {
	args := taskUseCaseMock.Called(sessionId, taskId)

	return args.Get(0).(dto.TaskDto), args.Error(1)
}

func (taskUseCaseMock *TaskUseCaseMock) GetAll(sessionId string) ([]dto.TaskDto, error) {
	args := taskUseCaseMock.Called(sessionId)

	return args.Get(0).([]dto.TaskDto), args.Error(1)
}

func (taskUseCaseMock *TaskUseCaseMock) Complete(sessionId string, taskId uint) (dto.TaskDto, error) {
	args := taskUseCaseMock.Called(sessionId, taskId)

	return args.Get(0).(dto.TaskDto), args.Error(1)
}

func (taskUseCaseMock *TaskUseCaseMock) Update(updateTaskDto dto.UpdateTaskDto) (dto.TaskDto, error) {
	args := taskUseCaseMock.Called(updateTaskDto)

	return args.Get(0).(dto.TaskDto), args.Error(1)
}

func (taskUseCaseMock *TaskUseCaseMock) Delete(sessionId string, taskId uint) error {
	args := taskUseCaseMock.Called(sessionId, taskId)

	return args.Error(0)
}
