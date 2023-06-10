package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	ThreadID int    `json:"thread_id" form:"thread_id"`
	Thread   Thread `json:"thread"`
	Comment  string `json:"comment" form:"comment"`
}
