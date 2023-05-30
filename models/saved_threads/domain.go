package savedthreads

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
	User_id int
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, keyword string) (Domain, error)
	Create(ctx context.Context, savedthreadDomain *Domain) (Domain, error)
	Update(ctx context.Context, savedthreadDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	GetByName(ctx context.Context, keyword string) (Domain, error)
	Create(ctx context.Context, savedthreadDomain *Domain) (Domain, error)
	Update(ctx context.Context, savedthreadDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) (Domain, error)
	ForceDelete(ctx context.Context, id string) error
}

