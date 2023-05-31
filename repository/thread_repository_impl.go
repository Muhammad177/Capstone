package repository

import (
	"context"
	"thread/model"

	"gorm.io/gorm"
)

type threadRepositoryImpl struct {
	db *gorm.DB
}

func (t *threadRepositoryImpl) CreateThread(thread *model.Thread, ctx context.Context) error {
	err := t.db.Save(thread).Error
	return err
}

// Deletethread implements threadRepository
func (t *threadRepositoryImpl) DeleteThread(thread *model.Thread, ctx context.Context) error {
	res := t.db.Delete(thread)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// FindAllthreads implements threadRepository
func (t *threadRepositoryImpl) FindAllThreads(ctx context.Context) ([]model.Thread, error) {
	var threads []model.Thread

	res := t.db.Find(&threads)
	if res.RowsAffected == 0 {
		return threads, gorm.ErrRecordNotFound
	}
	return threads, res.Error
}

// FindthreadBy implements threadRepository
func (t *threadRepositoryImpl) FindThreadBy(thread *model.Thread, ctx context.Context) ([]model.Thread, error) {
	var threads []model.Thread

	res := t.db.Where(thread).Find(&threads)
	if res.RowsAffected == 0 {
		return threads, gorm.ErrRecordNotFound
	}
	return threads, res.Error
}

// FindthreadById implements threadRepository
func (t *threadRepositoryImpl) FindThreadById(thread *model.Thread, ctx context.Context) error {
	err := t.db.First(thread).Error
	return err
}

// Updatethread implements threadRepository
func (t *threadRepositoryImpl) UpdateThread(thread *model.Thread, ctx context.Context) error {
	err := t.db.Model(thread).Updates(model.Thread{
		Title:     thread.Title,
		Topic:     thread.Topic,
		Content:   thread.Content,
		UpdatedAt: thread.UpdatedAt,
	}).Error

	return err
}

func NewThreadRepository(db *gorm.DB) ThreadRepository {
	return &threadRepositoryImpl{
		db: db,
	}
}
