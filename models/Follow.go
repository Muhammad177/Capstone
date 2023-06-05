package models

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Topic   string `json:"topic" form:"topic"`
	Content string `json:"content" form:"content"`
	File    string `json:"file" form:"file"`
	UserID  int    `json:"user_id" form:"user_id"`
	User    User   `json:"user"`
}
