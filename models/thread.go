package models

import (
	"github.com/jinzhu/gorm"
)

type Thread struct {
	gorm.Model
	Title    string    `json:"title" form:"title"`
	Topic    string    `json:"topic" form:"topic"`
	Content  string    `json:"content" form:"content"`
	File     string    `json:"file" form:"file"`
	UserID   int       `json:"user_id" form:"user_id"`
	User     User      `json:"user"`
	Comments []Comment `json:"comment"`
}

type ThreadResponse struct {
	Title    string            `json:"title" form:"title"`
	Topic    string            `json:"topic" form:"topic"`
	Comments []CommentResponse `json:"comments"`
}

func ConvertThreadToThreadResponse(thread *Thread) ThreadResponse {
	comments := make([]CommentResponse, len(thread.Comments))
	for i, comment := range thread.Comments {
		comments[i] = ConvertCommentToCommentResponse(&comment)
	}

	return ThreadResponse{
		Title:    thread.Title,
		Topic:    thread.Topic,
		Comments: comments,
	}
}
