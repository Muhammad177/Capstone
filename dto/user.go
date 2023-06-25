package dto

import (
	"Capstone/models"
)

type AllUser struct {
	ID        uint
	Username  string `json:"username" `
	Email     string `json:"email" `
	Image_url string `json:"image_url"`
	Age       int    `json:"age"`
	Bio       string `json:"bio"`
}

type DetailUserResponse struct {
	ID        uint
	Username  string            `json:"username" `
	Email     string            `json:"email" `
	Image_url string            `json:"image_url"`
	Age       int               `json:"age"`
	Bio       string            `json:"bio"`
	Followed  FollowedResponses `json:"followed"`
	Following int               `json:"following"`
	Follower  int               `json:"follower"`
}

func NewDetailUserResponse(user models.User, following int64, follower int64) *DetailUserResponse {
	return &DetailUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Image_url: user.Image_url,
		Age:       user.Age,
		Bio:       user.Bio,
		Followed:  *NewFollowedResponses(user.Followed),
		Following: int(following),
		Follower:  int(follower),
	}
}

func ConvertUserToAllUser(user models.User) *AllUser {
	return &AllUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Image_url: user.Image_url,
		Age:       user.Age,
		Bio:       user.Bio,
	}
}

type FollowUserRequest struct {
	ID int `json:"user_id", validate:"required"`
}

type FollowedResponse struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
}

func NewFollowedResponse(followed *models.User) *FollowedResponse {
	return &FollowedResponse{
		ID:       int(followed.ID),
		Username: followed.Username,
	}
}

type FollowedResponses []FollowedResponse

func NewFollowedResponses(data []models.User) *FollowedResponses {
	result := FollowedResponses{}

	for _, each := range data {
		result = append(result, *NewFollowedResponse(&each))
	}

	return &result
}
