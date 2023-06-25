package database

import (
	"Capstone/models"
	"context"
)

func FollowUser(ctx context.Context, user_id int, toBeFollow int) error {

	err := DB.WithContext(ctx).Create(&models.UserFollowedAssoc{UserID: uint(user_id), FollowedID: uint(toBeFollow)}).Error
	if err != nil {
		return err
	}

	return nil
}

func UnFollowUser(ctx context.Context, user_id int, followed_id int) error {

	err := DB.WithContext(ctx).Unscoped().Delete(&models.UserFollowedAssoc{UserID: uint(user_id), FollowedID: uint(followed_id)}).Error
	if err != nil {
		return err
	}

	return nil
}
