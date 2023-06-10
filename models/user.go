package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string   `json:"username" form:"username"`
	Email    string   `json:"email" form:"email"`
	Password string   `json:"password" form:"password"`
	Photo    string   `json:"photo" form:"photo"`
	Bio      string   `json:"bio" form:"bio"`
	Role     string   `json:"role" form:"role"`
	Threads  []Thread `json:"threads"`
	Follows  []Follow `json:"follows"`
}
type AllUser struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Photo    string `json:"photo" form:"photo"`
	Bio      string `json:"bio" form:"bio"`
}



func ConvertUserToAllUser(user *User) AllUser {
	return AllUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Photo:    user.Photo,
		Bio:      user.Bio,
	}
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
