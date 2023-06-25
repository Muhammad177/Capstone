package database

import (
	"Capstone/models"
	"context"
)

func GetMute(ctx context.Context) ([]models.Mute, error) {
	var mute []models.Mute

	err := DB.WithContext(ctx).Preload("User").Find(&mute).Error
	if err != nil {
		return nil, err
	}

	return mute, nil
}

func CreateMute(ctx context.Context, Mute models.Mute) (models.Mute, error) {
	err := DB.WithContext(ctx).Create(&Mute).Error
	if err != nil {
		return models.Mute{}, err
	}

	// Preload user data for the created Mute
	err = DB.WithContext(ctx).Preload("User").First(&Mute).Error
	if err != nil {
		return models.Mute{}, err
	}

	return Mute, nil
}

func DeleteMutes(ctx context.Context, id int) error {
	var Mute models.Mute

	result := DB.WithContext(ctx).Where("user_id = ?", id).Delete(&Mute)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}
