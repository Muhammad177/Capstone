package database

import (
	"Capstone/models"
	"context"
)

func CreateComment(ctx context.Context, Comment models.Comment) (models.Comment, error) {
	err := DB.WithContext(ctx).Create(&Comment).Error
	if err != nil {
		return models.Comment{}, err
	}

	// Preload user data for the created Comment
	err = DB.WithContext(ctx).Preload("User").Preload("Thread").First(&Comment).Error
	if err != nil {
		return models.Comment{}, err
	}

	return Comment, nil
}

func DeleteComments(ctx context.Context, commentID int, userID int) error {
	var Comment models.Comment

	result := DB.WithContext(ctx).Where("id = ? AND user_id = ?", commentID, userID).First(&Comment)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func UpdateComments(ctx context.Context, userID int, id int, Comment models.Comment) (interface{}, error) {
	var users models.User
	result := DB.WithContext(ctx).Where("id = ?", id).First(&users)
	if result.Error != nil {
		return Comment, result.Error
	}
	result = DB.WithContext(ctx).Model(&models.Comment{}).Where("id = ?", id).Updates(Comment)
	if result.Error != nil {
		return Comment, result.Error
	}

	if result.RowsAffected == 0 {
		return Comment, ErrIDNotFound
	}

	return Comment, nil
}

func GetComments(ctx context.Context, id int) ([]models.Comment, error) {

	var comment []models.Comment

	err := DB.WithContext(ctx).Preload("User").Preload("Thread").Where("thread_id = ?", id).Find(&comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func GetCommentID(ctx context.Context, id int) (models.Comment, error) {
	var comment models.Comment

	err := DB.WithContext(ctx).Preload("User").Preload("Thread").Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}
