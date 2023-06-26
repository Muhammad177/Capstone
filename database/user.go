package database

import (
	"Capstone/models"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetUsersByID(ctx context.Context, id int) (savedUser models.User, followingCount int64, followerCount int64, err error) {

	err = DB.WithContext(ctx).Preload("User").Where("id = ? ", id).First(&savedUser).Error
	if err != nil {
		return
	}
	err = DB.WithContext(ctx).Model(&models.User{Model: gorm.Model{ID: uint(id)}}).Preload("Followed").First(&savedUser).Error
	if err != nil {
		return
	}

	err = DB.WithContext(ctx).Table("user_followed_assocs").Select("count(*)").Where("user_id = ?", savedUser.ID).Scan(&followingCount).Error
	if err != nil {
		return
	}

	err = DB.WithContext(ctx).Table("user_followed_assocs").Select("count(*)").Where("followed_id = ?", savedUser.ID).Scan(&followerCount).Error
	if err != nil {
		return
	}
	fmt.Println(followerCount, followingCount)

	return
}

func UpdateUser(ctx context.Context, id int) error {
	var user models.User
	result := DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func DeleteUser(ctx context.Context, id int) error {
	var user models.User

	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrIDNotFound
	}

	return nil
}

func CreateUser(ctx context.Context, user models.User) (models.User, error) {
	var countSavedUser int64
	err := DB.Where("email = ?", user.Email).Count(&countSavedUser).Error
	if err != nil {
		return models.User{}, err
	}
	if countSavedUser != 0 {
		return models.User{}, errors.New("Email already exists")
	}

	err = DB.Save(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var userByEmail models.User

	err := DB.Where("email = ?", email).First(&userByEmail).Error
	if err != nil {
		return models.User{}, err
	}

	return userByEmail, nil
}
