package dto

import (
	"time"
)

type TaskDto struct {
	SessionId   string
	Id          uint
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateTaskDto struct {
	SessionId   string
	Title       string
	Description string
}

type UpdateTaskDto struct {
	SessionId   string
	Id          uint
	Title       string
	Description string
}
