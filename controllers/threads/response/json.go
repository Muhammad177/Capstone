package response

import (
	"Capstone/models/threads"
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title	string	`json:"title"`
	Topic   string	`json:"topic"`
	Content string	`json:"content"`
	File    string	`json:"file"`
	Role	string	`json:"role"`
	User_id int	`json:"user_id" gorm:"foreignKey"`
}

func FromDomain(domain threads.Domain) Thread {
	return Thread{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Title:	domain.Title,
		Topic:   domain.Topic,
		Content: domain.Content,
		File:    domain.File,
		Role:	domain.Role,
		User_id: domain.User_id,
	}
}
