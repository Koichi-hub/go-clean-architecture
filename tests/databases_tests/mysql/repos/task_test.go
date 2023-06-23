package repos

import (
	"go-clean-architecture/databases/mysql/repos"
	"go-clean-architecture/entities"
	"go-clean-architecture/usecases/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type testCase struct {
		name     string
		args     entities.Task
		expected entities.Task
	}

	tests := []testCase{
		{
			name: "Creating task",
			args: entities.Task{
				SessionId:   "1",
				Title:       "task 1",
				Description: "description for task 1",
			},
			expected: entities.Task{
				SessionId:   "1",
				Id:          1,
				Title:       "task 1",
				Description: "description for task 1",
				Completed:   false,
			},
		},
	}

	var taskRepo interfaces.TaskRepo = repos.NewTaskRepo()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received, err := taskRepo.Create(test.args)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, received)
		})
	}
}

func TestGetById(t *testing.T) {

}

func TestGetAll(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
