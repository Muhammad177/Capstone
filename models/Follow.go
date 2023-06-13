package models

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	ThreadID int    `json:"thread_id" form:"thread_id"`
	Thread   Thread `json:"thread"`
}
