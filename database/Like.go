package database

import (
	"Capstone/models"
	"context"
)

func CreateLike(ctx context.Context, Like models.Like) (models.Like, error) {
	err := DB.WithContext(ctx).Create(&Like).Error
	if err != nil {
		return models.Like{}, err
	}

	// Preload user data for the created Like
	err = DB.WithContext(ctx).Preload("Thread").First(&Like).Error
	if err != nil {
		return models.Like{}, err
	}

	return Like, nil
}

func DeleteLikes(ctx context.Context, id int) error {
	var Like models.Like

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&Like)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
func GetLikesByID(ctx context.Context, id int) ([]models.Like, error) {
	var Like []models.Like

	err := DB.WithContext(ctx).Preload("Thread").Where("user_id = ?", id).Find(&Like).Error
	if err != nil {
		return nil, err
	}

	return Like, nil
}
