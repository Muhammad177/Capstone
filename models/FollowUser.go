package models

import "github.com/jinzhu/gorm"

type FollowrUser struct {
	gorm.Model
	UserID int    `json:"user_id" form:"user_id"`
	User   User   `json:"user"`
	Status string `json:"status" form:"status"`
}
