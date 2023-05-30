package response

import (
	"Capstone/models/comments"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User_id int	`json:"user_id" gorm:"foreignKey"`
	Thread_id int	`json:"thread_id" gorm:"foreignKey"`
	Comment	string	`json:"comment"`
}

func FromDomain(domain comments.Domain) Comment {
	return Comment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		User_id: domain.User_id,
		Thread_id: domain.Thread_id,
		Comment: domain.Comment,
	}
}
