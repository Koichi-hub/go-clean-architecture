package controllers_tests

import (
	"bytes"
	"encoding/json"
	"go-clean-architecture/controllers"
	"go-clean-architecture/controllers/api_dto"
	"go-clean-architecture/tests/controllers_tests/usecases_mocks"
	"go-clean-architecture/usecases/dto"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Testing gin controllers, because gin has already tested http
// working with gin.Context

func TestCreate(t *testing.T) {
	type testCase struct {
		name     string
		args     api_dto.CreateTaskRequest
		expected api_dto.TaskResponse
	}

	tests := []testCase{
		{
			name: "Creating task",
			args: api_dto.CreateTaskRequest{
				Title:       "task 1",
				Description: "description for task 1",
			},
			expected: api_dto.TaskResponse{
				SessionId:   "1",
				Id:          1,
				Title:       "task 1",
				Description: "description for task 1",
				Completed:   false,
			},
		},
	}

	taskUseCaseMock := usecases_mocks.NewTaskUseCaseMock()
	taskUseCaseMock.On("Create", mock.AnythingOfType("dto.CreateTaskDto")).Return(dto.TaskDto{
		SessionId:   "1",
		Id:          1,
		Title:       "task 1",
		Description: "description for task 1",
		Completed:   false,
	}, nil).Once()

	taskController := controllers.NewTaskController(taskUseCaseMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(res)

			jsonValue, err := json.Marshal(test.args)
			assert.NoError(t, err)
			req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
			assert.NoError(t, err)

			ctx.Request = req
			ctx.Set("sessionId", "1")

			taskController.Create(ctx)
			assert.Equal(t, http.StatusCreated, res.Code)

			var received api_dto.TaskResponse
			json.Unmarshal(res.Body.Bytes(), &received)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestGetById(t *testing.T) {
	type args struct {
		taskId uint
	}

	type testCase struct {
		name string
		args
		expected api_dto.TaskResponse
	}

	tests := []testCase{
		{
			name: "Getting task by id",
			args: args{
				taskId: 1,
			},
			expected: api_dto.TaskResponse{
				SessionId:   "1",
				Id:          1,
				Title:       "task 1",
				Description: "description for task 1",
				Completed:   false,
			},
		},
	}

	taskUseCaseMock := usecases_mocks.NewTaskUseCaseMock()
	taskUseCaseMock.On("GetById", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(dto.TaskDto{
		SessionId:   "1",
		Id:          1,
		Title:       "task 1",
		Description: "description for task 1",
		Completed:   false,
	}, nil).Once()

	taskController := controllers.NewTaskController(taskUseCaseMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(res)

			req, err := http.NewRequest("GET", "/", nil)
			assert.NoError(t, err)

			ctx.Request = req
			ctx.Set("sessionId", "1")
			ctx.Params = []gin.Param{
				{
					Key:   "id",
					Value: strconv.FormatUint(uint64(test.args.taskId), 10),
				},
			}

			taskController.GetById(ctx)
			assert.Equal(t, http.StatusOK, res.Code)

			var received api_dto.TaskResponse
			json.Unmarshal(res.Body.Bytes(), &received)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestGetAll(t *testing.T) {
	type testCase struct {
		name     string
		expected []api_dto.TaskResponse
	}

	tests := []testCase{
		{
			name: "Getting all tasks",
			expected: []api_dto.TaskResponse{
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

	taskUseCaseMock := usecases_mocks.NewTaskUseCaseMock()
	taskUseCaseMock.On("GetAll", mock.AnythingOfType("string")).Return([]dto.TaskDto{
		{
			SessionId:   "1",
			Id:          1,
			Title:       "task 1",
			Description: "description for task 1",
			Completed:   false,
		},
	}, nil).Once()

	taskController := controllers.NewTaskController(taskUseCaseMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(res)

			req, err := http.NewRequest("GET", "/", nil)
			assert.NoError(t, err)

			ctx.Request = req
			ctx.Set("sessionId", "1")

			taskController.GetAll(ctx)
			assert.Equal(t, http.StatusOK, res.Code)

			var received []api_dto.TaskResponse
			json.Unmarshal(res.Body.Bytes(), &received)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestComplete(t *testing.T) {
	type args struct {
		taskId uint
	}

	type testCase struct {
		name string
		args
		expected api_dto.TaskResponse
	}

	tests := []testCase{
		{
			name: "Completing task",
			args: args{
				taskId: 1,
			},
			expected: api_dto.TaskResponse{
				SessionId:   "1",
				Id:          1,
				Title:       "task 1",
				Description: "description for task 1",
				Completed:   true,
			},
		},
	}

	taskUseCaseMock := usecases_mocks.NewTaskUseCaseMock()
	taskUseCaseMock.On("Complete", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(dto.TaskDto{
		SessionId:   "1",
		Id:          1,
		Title:       "task 1",
		Description: "description for task 1",
		Completed:   true,
	}, nil).Once()

	taskController := controllers.NewTaskController(taskUseCaseMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(res)

			req, err := http.NewRequest("PATCH", "/complete", nil)
			assert.NoError(t, err)

			ctx.Request = req
			ctx.Set("sessionId", "1")
			ctx.Params = []gin.Param{
				{
					Key:   "id",
					Value: strconv.FormatUint(uint64(test.args.taskId), 10),
				},
			}

			taskController.Complete(ctx)
			assert.Equal(t, http.StatusOK, res.Code)

			var received api_dto.TaskResponse
			json.Unmarshal(res.Body.Bytes(), &received)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		taskId            uint
		updateTaskRequest api_dto.UpdateTaskRequest
	}

	type testCase struct {
		name string
		args
		expected api_dto.TaskResponse
	}

	tests := []testCase{
		{
			name: "Updating task",
			args: args{
				taskId: 1,
				updateTaskRequest: api_dto.UpdateTaskRequest{
					Title:       "task 2",
					Description: "description for task 2",
				},
			},
			expected: api_dto.TaskResponse{
				SessionId:   "1",
				Id:          1,
				Title:       "task 2",
				Description: "description for task 2",
				Completed:   false,
			},
		},
	}

	taskUseCaseMock := usecases_mocks.NewTaskUseCaseMock()
	taskUseCaseMock.On("Update", mock.AnythingOfType("dto.UpdateTaskDto")).Return(dto.TaskDto{
		SessionId:   "1",
		Id:          1,
		Title:       "task 2",
		Description: "description for task 2",
		Completed:   false,
	}, nil).Once()

	taskController := controllers.NewTaskController(taskUseCaseMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(res)

			jsonValue, err := json.Marshal(test.args.updateTaskRequest)
			assert.NoError(t, err)
			req, err := http.NewRequest("PUT", "/", bytes.NewBuffer(jsonValue))
			assert.NoError(t, err)

			ctx.Request = req
			ctx.Set("sessionId", "1")
			ctx.Params = []gin.Param{
				{
					Key:   "id",
					Value: strconv.FormatUint(uint64(test.args.taskId), 10),
				},
			}

			taskController.Update(ctx)
			assert.Equal(t, http.StatusOK, res.Code)

			var received api_dto.TaskResponse
			json.Unmarshal(res.Body.Bytes(), &received)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		taskId uint
	}

	type testCase struct {
		name string
		args
	}

	tests := []testCase{
		{
			name: "Deleting task",
			args: args{
				taskId: 1,
			},
		},
	}

	taskUseCaseMock := usecases_mocks.NewTaskUseCaseMock()
	taskUseCaseMock.On("Delete", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil).Once()

	taskController := controllers.NewTaskController(taskUseCaseMock)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(res)

			req, err := http.NewRequest("DELETE", "/", nil)
			assert.NoError(t, err)

			ctx.Request = req
			ctx.Set("sessionId", "1")
			ctx.Params = []gin.Param{
				{
					Key:   "id",
					Value: strconv.FormatUint(uint64(test.args.taskId), 10),
				},
			}

			taskController.Delete(ctx)
			assert.Equal(t, http.StatusOK, res.Code)
		})
	}
}
