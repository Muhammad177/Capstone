package database

import (
	"Capstone/models"
	"context"

	"gorm.io/gorm"
)

func LikeThread(ctx context.Context, user_id int, thread_id int) error {

	err := DB.WithContext(ctx).Create(&models.ThreadLikeAssoc{UserID: uint(user_id), ThreadID: uint(thread_id)}).Error
	if err != nil {
		return err
	}

	err = DB.WithContext(ctx).Model(&models.Thread{}).Where("id = ?", thread_id).Update("like_count", gorm.Expr("like_count+1")).Error
	if err != nil {
		return err
	}

	return nil
}

func UnlikeThread(ctx context.Context, user_id int, thread_id int) error {

	err := DB.WithContext(ctx).Unscoped().Delete(&models.ThreadLikeAssoc{UserID: uint(user_id), ThreadID: uint(thread_id)}).Error
	if err != nil {
		return err
	}

	err = DB.WithContext(ctx).Model(&models.Thread{}).Where("id = ?", thread_id).Update("like_count", gorm.Expr("like_count-1")).Error
	if err != nil {
		return err
	}

	return nil
}
