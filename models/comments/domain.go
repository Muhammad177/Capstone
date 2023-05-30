package comments

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
	Thread_id int
	User_id	int
	Comment	string
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, keyword string) (Domain, error)
	Create(ctx context.Context, commentDomain *Domain) (Domain, error)
	Update(ctx context.Context, commentDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, keyword string) (Domain, error)
	Create(ctx context.Context, commentDomain *Domain) (Domain, error)
	Update(ctx context.Context, commentDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

