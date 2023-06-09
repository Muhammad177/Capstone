package models

type Follow struct {
	FollowerID int  `json:"follow"`
	UserID     int  `json:"user_id" form:"user_id"`
	User       User `json:"user"`
}
