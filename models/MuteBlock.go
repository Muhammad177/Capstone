package models

import "github.com/jinzhu/gorm"

type Mute struct {
	gorm.Model
	UserID float64 `json:"user_id" form:"user_id"`
	User   User    `json:"User"`
	Status string  `json:"status" form:"status"`
}
