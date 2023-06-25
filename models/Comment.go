package models

import (
	"github.com/jinzhu/gorm"
)


type Comment struct {
	gorm.Model
	Comment  string `json:"comment" form:"comment" validate:"required"`
	ThreadID int    `json:"thread_id" form:"thread_id" validate:"required"`
	Thread   Thread `json:"thread"`
	UserID   int    `json:"user_id" form:"user_id" validate:"required"`
	User     User   `json:"user"`
}
type CommentResponse struct {
	ID       uint          `gorm:"primary_key"`
	Comment  string        `json:"comment" form:"comment"`
	ThreadID int           `json:"thread_id" form:"thread_id"`
	Thread   Thread        `json:"thread"`
	UserID   int           `json:"user_id" form:"user_id"`
	Username AllUserSearch `json:"user"`
}

func ConvertCommentToCommentResponse(comment *Comment) CommentResponse {
	return CommentResponse{
		ID:       comment.ID,
		Comment:  comment.Comment,
		ThreadID: comment.ThreadID,
		Thread:   comment.Thread,
		UserID:   comment.UserID,
		Username: ConvertAllUserSearch(&comment.User),
	}
}
