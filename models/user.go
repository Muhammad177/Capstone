package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string   `json:"username" form:"username"`
	Email      string   `json:"email" form:"email"`
	Password   string   `json:"password" form:"password"`
	Photo      string   `json:"photo" form:"photo"`
	Bio        string   `json:"bio" form:"bio"`
	Role       string   `json:"role" form:"role"`
	Threads    []Thread `json:"threads"`
	Bookmarked []Thread `gorm:"many2many: thread_user_assoc"`
	Follows    []Follow `json:"follows"`
}
type UserResponse struct {
	ID    int    `json:"id" form:"name"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
type AdminResponse struct {
	ID       int    `json:"id" form:"name"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
