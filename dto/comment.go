package dto

import (
	"Capstone/models"
	"time"
)

type GetCommentResponse struct {
	ID       int            `json:"id"`
	Comment  string         `json:"comment"`
	Author   AuthorResponse `json:"author"`
	CreateAt time.Time      `json:"createdAt"`
	UpdateAt time.Time      `json:"updateAt"`
}

func NewGetCommentResponse(data models.Comment) *GetCommentResponse {
	return &GetCommentResponse{
		ID:       int(data.ID),
		Comment:  data.Comment,
		Author:   *NewAuthorResponse(data.User),
		CreateAt: data.CreatedAt,
		UpdateAt: data.UpdatedAt,
	}
}

type AuthorsResponse struct {
	ID       int    `json:"user_ID"`
	Username string `json:"username"`
}

func NewAuthorsResponse(author models.User) *AuthorsResponse {
	return &AuthorsResponse{
		ID:       int(author.ID),
		Username: author.Username,
	}
}

type GetCommentsResponse []GetCommentResponse

func NewGetCommentsResponse(data []models.Comment) *GetCommentsResponse {
	result := GetCommentsResponse{}

	for _, each := range data {
		result = append(result, *NewGetCommentResponse(each))
	}

	return &result
}
