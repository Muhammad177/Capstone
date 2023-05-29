package users

import (
	"Capstone/models/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username  string         `json:"username" gorm:"unique"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Photo     string         `json:"photo"`
	Bio       string         `json:"bio"`
	Role      string         `json:"role"`
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Username:  rec.Username,
		Email:     rec.Email,
		Password:  rec.Password,
		Photo:     rec.Photo,
		Bio:       rec.Bio,
		Role:      rec.Role,
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
		Photo:     domain.Photo,
		Bio:       domain.Bio,
		Role:      domain.Role,
	}
}
