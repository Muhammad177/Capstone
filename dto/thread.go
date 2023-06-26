package dto

import (
	"Capstone/models"
	"time"
)

type DetailThreadResponse struct {
	ID           uint
	Title        string              `json:"title"`
	Content      string              `json:"content"`
	File         string              `json:"file"`
	Author       AuthorResponse      `json:"author"`
	CommentCount int                 `json:"commentCount`
	Comment      GetCommentsResponse `json:"comment"`
	LikeUser     LikeUsersResponse   `json:"likeUser"`
	Like         int                 `json:"like"`
}

func NewDetailThreadResponse(thread *models.Thread) *DetailThreadResponse {
	return &DetailThreadResponse{
		ID:           thread.ID,
		Title:        thread.Title,
		Content:      thread.Content,
		File:         thread.File,
		Author:       *NewAuthorResponse(thread.User),
		Comment:      *NewGetCommentsResponse(thread.Comments),
		CommentCount: thread.CommentCount,
		LikeUser:     *NewLikeUsersResponse(thread.Like),
		Like:         thread.LikeCount,
	}
}

type GetAllThreadResponse struct {
	ID       uint
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	File     string            `json:"file"`
	Comment  int               `json:"comment"`
	Author   AuthorResponse    `json:"author"`
	LikeUser LikeUsersResponse `json:"likeUser"`
	Like     int               `json:"like"`
}

func NewGetAllThreadResponse(thread *models.Thread) *GetAllThreadResponse {
	return &GetAllThreadResponse{
		ID:       thread.ID,
		Title:    thread.Title,
		Content:  thread.Content,
		File:     thread.File,
		Author:   *NewAuthorResponse(thread.User),
		Comment:  thread.CommentCount,
		LikeUser: *NewLikeUsersResponse(thread.Like),
		Like:     thread.LikeCount,
	}
}

type GetAllThreadResponses []GetAllThreadResponse

func NewGetAllThreadResponses(data []models.Thread) *GetAllThreadResponses {
	result := GetAllThreadResponses{}

	for _, each := range data {
		result = append(result, *NewGetAllThreadResponse(&each))
	}
	return &result
}

type LikeUserResponse struct {
	ID       uint   `json:"user_id"`
	Username string `json:"username"`
}

func NewLikeUserResponse(user *models.User) *LikeUserResponse {
	return &LikeUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}

type LikeUsersResponse []LikeUserResponse

func NewLikeUsersResponse(data []models.User) *LikeUsersResponse {
	result := LikeUsersResponse{}

	for _, each := range data {
		result = append(result, *NewLikeUserResponse(&each))
	}

	return &result
}

type CommentResponse struct {
	ID       int            `json:"id"`
	Comment  string         `json:"comment"`
	Author   AuthorResponse `json:"author"`
	CreateAt time.Time      `json:"createdAt"`
	UpdateAt time.Time      `json:"updateAt"`
}

func NewCommentResponse(data models.Comment) *CommentResponse {
	return &CommentResponse{
		ID:       int(data.ID),
		Comment:  data.Comment,
		Author:   *NewAuthorResponse(data.User),
		CreateAt: data.CreatedAt,
		UpdateAt: data.UpdatedAt,
	}
}

type CommmentAuthorResponse struct {
	ID       int    `json:"user_ID"`
	Username string `json:"username"`
}

func NewCommmentAuthorResponse(author models.User) *AuthorsResponse {
	return &AuthorsResponse{
		ID:       int(author.ID),
		Username: author.Username,
	}
}

type CommentsResponse []CommentResponse

func NewCommentsResponse(data []models.Comment) *CommentsResponse {
	result := CommentsResponse{}

	for _, each := range data {
		result = append(result, *NewCommentResponse(each))
	}

	return &result
}
