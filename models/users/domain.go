package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Photo     string
	Bio       string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, error)
}

type Repository interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
}
