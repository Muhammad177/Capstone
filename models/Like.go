package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	UserID   int    `json:"user_id" form:"user_id"`
	ThreadID int    `json:"thread_id" form:"thread_id" validate:"required"`
	Thread   Thread `json:"thread"`
	User     User   `json:"user"`
}
