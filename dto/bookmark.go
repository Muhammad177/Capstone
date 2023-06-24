package dto

import (
	"Capstone/models"
	"time"
)

// "Capstone/models"

// "gorm.io/gorm"

type GetBookmarkResponse struct {
	ThreadID  int            `json:"thread_id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Author    AuthorResponse `json:"author"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updatedAt"`
}

func NewGetBookmarkResponse(data models.Thread) *GetBookmarkResponse {
	return &GetBookmarkResponse{
		ThreadID:  int(data.ID),
		Title:     data.Title,
		Content:   data.Content,
		Author:    *NewAuthorResponse(data.User),
		CreatedAt: data.CreatedAt,
		UpdateAt:  data.UpdatedAt,
	}
}

type AuthorResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewAuthorResponse(author models.User) *AuthorResponse {
	return &AuthorResponse{
		Username: author.Username,
		Email:    author.Email,
	}
}

type GetBookmarksResponse []GetBookmarkResponse

func NewGetBookmarksResponse(data models.User) *GetBookmarksResponse {
	result := GetBookmarksResponse{}

	for _, each := range data.Bookmarked {
		result = append(result, *NewGetBookmarkResponse(each))
	}

	return &result
}
