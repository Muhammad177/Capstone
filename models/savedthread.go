package models

import (
	"github.com/jinzhu/gorm"
)

type Savedthread struct {
	gorm.Model
	ID           uint           `json:"id" gorm:"primaryKey"`
	User_id	 	 uint 			`json:"user" gorm:"foreignkey:User_id"`
	Thread_id	 uint 			`json:"thread" gorm:"foreignkey:Thread_id"`
}
