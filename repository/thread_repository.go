package repository

import (
	"context"
	"thread/model"
)

type ThreadRepository interface {
	CreateThread(thread *model.Thread, ctx context.Context) error
	UpdateThread(thread *model.Thread, ctx context.Context) error
	DeleteThread(thread *model.Thread, ctx context.Context) error
	FindAllThreads(ctx context.Context) ([]model.Thread, error)
	FindThreadById(thread *model.Thread, ctx context.Context) error
	FindThreadBy(thread *model.Thread, ctx context.Context) ([]model.Thread, error)
}
