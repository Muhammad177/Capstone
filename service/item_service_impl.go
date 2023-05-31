package service

import (
	"context"
	"thread/config"
	"thread/model"
	"thread/repository"
	"time"

	"github.com/google/uuid"
)

type threadServiceImpl struct {
	repo repository.ThreadRepository
}

// CreateItem implements itemService
func (t *threadServiceImpl) CreateThread(thread *model.Thread, ctx context.Context) error {
	thread.ID = uuid.New().String()
	loc, _ := time.LoadLocation(config.Cfg.TIME_LOC)
	thread.CreatedAt = time.Now().In(loc)
	thread.UpdatedAt = time.Now().In(loc)
	err := t.repo.CreateThread(thread, ctx)
	return err
}

// DeleteItem implements itemService
func (t *threadServiceImpl) DeleteThread(thread *model.Thread, ctx context.Context) error {
	err := t.repo.DeleteThread(thread, ctx)
	return err
}

// GetAllItems implements itemService
func (t *threadServiceImpl) GetAllThreads(keyword string, ctx context.Context) ([]model.Thread, error) {
	if keyword != "" {
		var thread model.Thread
		thread.Title = keyword
		res, err := t.repo.FindThreadBy(&thread, ctx)
		return res, err
	}
	res, err := t.repo.FindAllThreads(ctx)
	return res, err
}

// GetItemBy implements itemService
func (t *threadServiceImpl) GetThreadBy(thread *model.Thread, ctx context.Context) ([]model.Thread, error) {
	res, err := t.repo.FindThreadBy(thread, ctx)
	return res, err
}

// GetItemById implements itemService
func (t *threadServiceImpl) GetThreadById(thread *model.Thread, ctx context.Context) error {
	err := t.repo.FindThreadById(thread, ctx)
	return err
}

// UpdateItem implements itemService
func (t *threadServiceImpl) UpdateThread(thread *model.Thread, ctx context.Context) error {
	thread.UpdatedAt = time.Now()

	err := t.repo.UpdateThread(thread, ctx)
	return err
}

func NewThreadService(repo repository.ThreadRepository) ThreadService {
	return &threadServiceImpl{
		repo: repo,
	}
}
