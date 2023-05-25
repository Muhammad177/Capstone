package models

import (
	"github.com/jinzhu/gorm"
)

type Thread struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Topic   string `json:"topic" form:"topic"`
	Content string `json:"content" form:"content"`
	File    string `json:"file" form:"file"`
	User_id int    `json:user_id`
	User    User   `json:user`
}
