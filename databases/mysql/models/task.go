package models

import "time"

type TaskModel struct {
	SessionId   string    `gorm:"primaryKey;autoIncrement:false;column:session_id"`
	Id          uint      `gorm:"primaryKey;column:id"`
	Title       string    `gorm:"size:50;column:title"`
	Description string    `gorm:"size:1000;column:description"`
	Completed   bool      `gorm:"column:completed"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
