package repos

import (
	"database/sql/driver"
	"go-clean-architecture/databases/mysql/repos"
	"go-clean-architecture/entities"
	"go-clean-architecture/usecases/interfaces"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	type testCase struct {
		name     string
		args     entities.Task
		expected uint
	}

	tests := []testCase{
		{
			name: "Creating task",
			args: entities.Task{
				SessionId:   "1",
				Title:       "task 1",
				Description: "description for task 1",
			},
			expected: 1,
		},
	}

	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer mockDb.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}

	var taskRepo interfaces.TaskRepo = repos.NewTaskRepo(db)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := "INSERT INTO `task_models` (`session_id`,`title`,`description`,`completed`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)"
			queryArgs := []driver.Value{test.args.SessionId, test.args.Title, test.args.Description, test.args.Completed, sqlmock.AnyArg(), sqlmock.AnyArg()}

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(queryArgs...).
				WillReturnResult(sqlmock.NewResult(int64(test.expected), 1))
			mock.ExpectCommit()

			err := taskRepo.Create(test.args)

			assert.NoError(t, err)
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
