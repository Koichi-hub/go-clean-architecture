package models

import "time"

type TaskModel struct {
	SessionId   string `gorm:"primaryKey"`
	Id          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:50"`
	Description string `gorm:"size:1000"`
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
