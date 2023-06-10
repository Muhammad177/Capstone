package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	ThreadID int    `json:"thread_id" form:"thread_id"`
	Thread   Thread `json:"thread"`
	UserID   int       `json:"user_id" form:"user_id"`
	User     User      `json:"user"`
	Comment  string `json:"comment" form:"comment"`
}
