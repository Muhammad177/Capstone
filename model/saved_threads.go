package model

import (
	"time"

	"gorm.io/gorm"
)

type Saved_thread struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	UserID    int            `json:"user_id"`
	Users     User           `json: "users"`
	ThreadID  string         `json:"thread_id"`
	Thread    Thread         `json:"thread"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
