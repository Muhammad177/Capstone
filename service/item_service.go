package service

import (
	"context"
	"thread/model"
)

type ThreadService interface {
	GetAllThreads(keyword string, ctx context.Context) ([]model.Thread, error)
	CreateThread(thread *model.Thread, ctx context.Context) error
	UpdateThread(thread *model.Thread, ctx context.Context) error
	DeleteThread(thread *model.Thread, ctx context.Context) error
	GetThreadBy(thread *model.Thread, ctx context.Context) ([]model.Thread, error)
	GetThreadById(thread *model.Thread, ctx context.Context) error
}
