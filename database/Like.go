package database

import (
	"Capstone/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func CreateLikeThreads(ctx context.Context, user_id int, thread_id int) (models.User, error) {
	var thread models.Thread

	fmt.Println(thread_id, user_id)
	err := DB.WithContext(ctx).Where("id = ?", thread_id).First(&thread).Error
	if err != nil {
		return models.User{}, err
	}

	fmt.Print(thread)
	err = DB.WithContext(ctx).Model(&models.User{Model: gorm.Model{ID: uint(user_id)}}).Association("Liked").Append(&thread)
	if err != nil {
		return models.User{}, err
	}

	var lkThread models.User
	err = DB.WithContext(ctx).Preload("Liked").Find(&lkThread).Error
	if err != nil {
		return models.User{}, err
	}

	return lkThread, nil
}

func DeleteLikeThreads(ctx context.Context, user_id int, thread_id int) error {
	var thread models.Thread

	fmt.Println(thread_id, user_id)
	err := DB.WithContext(ctx).Where("id = ?", thread_id).First(&thread).Error
	if err != nil {
		return err
	}

	err = DB.WithContext(ctx).Model(&models.User{Model: gorm.Model{ID: uint(user_id)}}).Association("Liked").Delete(&thread)
	if err != nil {
		return err
	}

	return nil
}

func GetLikeThreads(ctx context.Context, user_id int) ([]models.Thread, error) {
	var threads []models.Thread

	err := DB.WithContext(ctx).Model(&models.User{Model: gorm.Model{ID: uint(user_id)}}).Association("Liked").Find(&threads)
	if err != nil {
		return nil, err
	}

	return threads, nil
}