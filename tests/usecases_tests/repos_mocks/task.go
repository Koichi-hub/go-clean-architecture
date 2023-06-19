package repos_mocks

import (
	"go-clean-architecture/entities"

	"github.com/stretchr/testify/mock"
)

type TaskRepoMock struct {
	mock.Mock
}

func NewTaskRepoMock() *TaskRepoMock {
	return &TaskRepoMock{}
}

func (taskRepoMock *TaskRepoMock) Save(task entities.Task) (entities.Task, error) {
	args := taskRepoMock.Called(task)

	return args.Get(0).(entities.Task), args.Error(1)
}

func (taskRepoMock *TaskRepoMock) GetById(sessionId string, taskId uint) (entities.Task, error) {
	args := taskRepoMock.Called(sessionId, taskId)

	return args.Get(0).(entities.Task), args.Error(1)
}

func (taskRepoMock *TaskRepoMock) GetAll(sessionId string) ([]entities.Task, error) {
	args := taskRepoMock.Called(sessionId)

	return args.Get(0).([]entities.Task), args.Error(1)
}

func (taskRepoMock *TaskRepoMock) Complete(sessionId string, taskId uint) (entities.Task, error) {
	args := taskRepoMock.Called(sessionId, taskId)

	return args.Get(0).(entities.Task), args.Error(1)
}

func (taskRepoMock *TaskRepoMock) Delete(sessionId string, taskId uint) error {
	args := taskRepoMock.Called(sessionId, taskId)

	return args.Error(0)
}
