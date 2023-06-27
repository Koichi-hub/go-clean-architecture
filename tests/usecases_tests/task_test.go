package usecases_tests

import (
	"go-clean-architecture/entities"
	"go-clean-architecture/tests/usecases_tests/repos_mocks"
	"go-clean-architecture/usecases"
	"go-clean-architecture/usecases/dto"
	"go-clean-architecture/usecases/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	type testCase struct {
		name     string
		args     dto.CreateTaskDto
		expected uint
	}

	tests := []testCase{
		{
			name: "Creating task",
			args: dto.CreateTaskDto{
				SessionId:   "1",
				Title:       "task 1",
				Description: "description for task 1",
			},
			expected: 1,
		},
	}

	taskRepoMock := repos_mocks.NewTaskRepoMock()
	taskRepoMock.On("Create", mock.AnythingOfType("entities.Task")).Return(1, nil).Once()

	var taskUseCase interfaces.TaskUseCase = usecases.NewTaskUseCase(taskRepoMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			taskId, err := taskUseCase.Create(test.args)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, taskId)
		})
	}
}

func TestGetById(t *testing.T) {
	type args struct {
		sessionId string
		taskId    uint
	}

	type testCase struct {
		name string
		args
		expected dto.TaskDto
	}

	tests := []testCase{
		{
			name: "Getting task by id",
			args: args{
				sessionId: "1",
				taskId:    1,
			},
			expected: dto.TaskDto{
				SessionId:   "1",
				Id:          1,
				Title:       "task 1",
				Description: "description for task 1",
				Completed:   false,
			},
		},
	}

	taskRepoMock := repos_mocks.NewTaskRepoMock()
	taskRepoMock.On("GetById", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(entities.Task{
		SessionId:   "1",
		Id:          1,
		Title:       "task 1",
		Description: "description for task 1",
		Completed:   false,
	}, nil).Once()

	var taskUseCase interfaces.TaskUseCase = usecases.NewTaskUseCase(taskRepoMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received, err := taskUseCase.GetById(test.args.sessionId, test.args.taskId)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestGetAll(t *testing.T) {
	type args struct {
		sessionId string
	}

	type testCase struct {
		name string
		args
		expected []dto.TaskDto
	}

	tests := []testCase{
		{
			name: "Getting all tasks",
			args: args{
				sessionId: "1",
			},
			expected: []dto.TaskDto{
				{
					SessionId:   "1",
					Id:          1,
					Title:       "task 1",
					Description: "description for task 1",
					Completed:   false,
				},
			},
		},
	}

	taskRepoMock := repos_mocks.NewTaskRepoMock()
	taskRepoMock.On("GetAll", mock.AnythingOfType("string")).Return([]entities.Task{
		{
			SessionId:   "1",
			Id:          1,
			Title:       "task 1",
			Description: "description for task 1",
			Completed:   false,
		},
	}, nil).Once()

	var taskUseCase interfaces.TaskUseCase = usecases.NewTaskUseCase(taskRepoMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received, err := taskUseCase.GetAll(test.args.sessionId)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestComplete(t *testing.T) {
	type args struct {
		sessionId string
		taskId    uint
	}

	type testCase struct {
		name string
		args
	}

	tests := []testCase{
		{
			name: "Completing task",
			args: args{
				sessionId: "1",
				taskId:    1,
			},
		},
	}

	taskRepoMock := repos_mocks.NewTaskRepoMock()
	taskRepoMock.On("GetById", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(entities.Task{
		SessionId:   "1",
		Id:          1,
		Title:       "task 1",
		Description: "description for task 1",
		Completed:   false,
	}, nil).Once()
	taskRepoMock.On("Update", mock.AnythingOfType("entities.Task")).Return(nil).Once()

	var taskUseCase interfaces.TaskUseCase = usecases.NewTaskUseCase(taskRepoMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := taskUseCase.Complete(test.args.sessionId, test.args.taskId)

			assert.NoError(t, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	type testCase struct {
		name string
		args dto.UpdateTaskDto
	}

	tests := []testCase{
		{
			name: "Updating task",
			args: dto.UpdateTaskDto{
				SessionId:   "1",
				Id:          1,
				Title:       "task 2",
				Description: "description for task 2",
			},
		},
	}

	taskRepoMock := repos_mocks.NewTaskRepoMock()
	taskRepoMock.On("Update", mock.AnythingOfType("entities.Task")).Return(nil).Once()

	var taskUseCase interfaces.TaskUseCase = usecases.NewTaskUseCase(taskRepoMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := taskUseCase.Update(test.args)

			assert.NoError(t, err)
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		sessionId string
		taskId    uint
	}

	type testCase struct {
		name string
		args
	}

	tests := []testCase{
		{
			name: "Deleting task",
			args: args{
				sessionId: "1",
				taskId:    1,
			},
		},
	}

	taskRepoMock := repos_mocks.NewTaskRepoMock()
	taskRepoMock.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil).Once()

	var taskUseCase interfaces.TaskUseCase = usecases.NewTaskUseCase(taskRepoMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := taskUseCase.Delete(test.args.sessionId, test.args.taskId)

			assert.NoError(t, err)
		})
	}
}
