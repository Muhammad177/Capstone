package model

import (
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Topic     string         `json:"topic"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
