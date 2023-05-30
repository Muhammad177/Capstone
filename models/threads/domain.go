package threads

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Title	string
	Topic   string
	Content string 
	File    string
	User_id int
	Role	string
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, keyword string) (Domain, error)
	Create(ctx context.Context, threadDomain *Domain) (Domain, error)
	Update(ctx context.Context, threadDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, keyword string) (Domain, error)
	Create(ctx context.Context, threadDomain *Domain) (Domain, error)
	Update(ctx context.Context, threadDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

