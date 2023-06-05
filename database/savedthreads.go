package database

import (
	"Capstone/models"
	"context"
)

func GetSavedthreads(ctx context.Context) ([]models.Savedthread, error) {
	var savedthread []models.Savedthread

	err := DB.WithContext(ctx).Find(&savedthread).Error
	if err != nil {
		return nil, err
	}

	return savedthread, nil
}

func GetSavedthreadsByID(ctx context.Context, id int) (models.Savedthread, error) {
	var savedthread models.Savedthread

	err := DB.WithContext(ctx).Where("id = ?", id).First(&savedthread).Error
	if err != nil {
		return models.Savedthread{}, err
	}

	return savedthread, nil
}

func CreateSavedthreads(ctx context.Context, savedthread models.Savedthread) (models.Savedthread, error) {
	err := DB.WithContext(ctx).Create(&savedthread).Error

	if err != nil {
		return models.Savedthread{}, err
	}

	return savedthread, nil
}

func DeleteSavedThreads(ctx context.Context, id int) error {
	var savedthread models.Savedthread

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&savedthread)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateSavedthreads(ctx context.Context, id int, savedthread models.Savedthread) (interface{}, error) {
	result := DB.WithContext(ctx).Model(&models.Savedthread{}).Where("id = ?", id).Updates(savedthread)
	if result.Error != nil {
		return savedthread, result.Error
	}

	if result.RowsAffected == 0 {
		return savedthread, ErrIDNotFound
	}

	return savedthread, nil
}
