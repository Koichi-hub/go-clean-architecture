package entities

import "time"

type Task struct {
	SessionId   string
	Id          uint
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
