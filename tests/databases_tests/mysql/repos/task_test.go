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
			query := "INSERT INTO `tasks` (`session_id`,`title`,`description`,`completed`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)"
			queryArgs := []driver.Value{test.args.SessionId, test.args.Title, test.args.Description, test.args.Completed, sqlmock.AnyArg(), sqlmock.AnyArg()}

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(queryArgs...).
				WillReturnResult(sqlmock.NewResult(int64(test.expected), 1))
			mock.ExpectCommit()

			taskId, err := taskRepo.Create(test.args)

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
		expected entities.Task
	}

	tests := []testCase{
		{
			name: "Getting task by id",
			args: args{
				sessionId: "1",
				taskId:    1,
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
			query := "SELECT * FROM `tasks` WHERE session_id = ? AND id = ? ORDER BY `tasks`.`id` LIMIT 1"

			mock.ExpectQuery(regexp.QuoteMeta(query)).
				WillReturnRows(
					sqlmock.NewRows([]string{"session_id", "id", "title", "description", "completed", "created_at", "updated_at"}).
						AddRow(test.expected.SessionId, test.expected.Id, test.expected.Title, test.expected.Description, test.expected.Completed, test.expected.CreatedAt, test.expected.UpdatedAt),
				)

			task, err := taskRepo.GetById(test.args.sessionId, test.args.taskId)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, task)
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
		expected []entities.Task
	}

	tests := []testCase{
		{
			name: "Getting all tasks",
			args: args{
				sessionId: "1",
			},
			expected: []entities.Task{
				{
					SessionId:   "1",
					Id:          1,
					Title:       "task 1",
					Description: "description for task 1",
					Completed:   false,
				},
				{
					SessionId:   "1",
					Id:          2,
					Title:       "task 2",
					Description: "description for task 2",
					Completed:   false,
				},
				{
					SessionId:   "1",
					Id:          3,
					Title:       "task 3",
					Description: "description for task 3",
					Completed:   true,
				},
			},
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
			query := "SELECT * FROM `tasks` WHERE session_id = ?"

			rows := sqlmock.NewRows([]string{"session_id", "id", "title", "description", "completed", "created_at", "updated_at"})
			for _, v := range test.expected {
				rows.AddRow(v.SessionId, v.Id, v.Title, v.Description, v.Completed, v.CreatedAt, v.UpdatedAt)
			}

			mock.ExpectQuery(regexp.QuoteMeta(query)).
				WillReturnRows(rows)

			tasks, err := taskRepo.GetAll(test.args.sessionId)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, tasks)
		})
	}
}

func TestUpdate(t *testing.T) {
	type testCase struct {
		name     string
		args     entities.Task
		expected uint
	}

	tests := []testCase{
		{
			name: "Updating task",
			args: entities.Task{
				SessionId:   "1",
				Id:          1,
				Title:       "updated title of task 1",
				Description: "updated description of task 1",
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
			query := "UPDATE `tasks` SET `title`=?,`description`=?,`completed`=?,`created_at`=?,`updated_at`=? WHERE `id` = ? AND `session_id` = ?"
			queryArgs := []driver.Value{test.args.Title, test.args.Description, test.args.Completed, sqlmock.AnyArg(), sqlmock.AnyArg(), test.args.Id, test.args.SessionId}

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(queryArgs...).
				WillReturnResult(sqlmock.NewResult(int64(test.expected), 1))
			mock.ExpectCommit()

			err := taskRepo.Update(test.args)

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
		expected uint
	}

	tests := []testCase{
		{
			name: "Deleting task",
			args: args{
				sessionId: "1",
				taskId:    1,
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
			query := "DELETE FROM `tasks` WHERE session_id = ? AND id = ?"
			queryArgs := []driver.Value{test.args.sessionId, test.args.taskId}

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(queryArgs...).
				WillReturnResult(sqlmock.NewResult(int64(test.expected), 1))
			mock.ExpectCommit()

			err := taskRepo.Delete(test.args.sessionId, test.args.taskId)

			assert.NoError(t, err)
		})
	}
}
