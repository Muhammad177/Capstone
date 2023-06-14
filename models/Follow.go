package models

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	ThreadID int    `json:"thread_id" form:"thread_id"`
	UserID   int    `json:"user_id" form:"user_id"`
	Thread   Thread `json:"thread"`
}
