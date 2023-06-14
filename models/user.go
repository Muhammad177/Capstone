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
	Age        int      `json:"age" form:"age"`
	Bio        string   `json:"bio" form:"bio"`
	Role       string   `json:"role" form:"role"`
	Threads    []Thread `json:"threads"`
	Bookmarked []Thread `gorm:"many2many: thread_user_assoc"`
}
type AllUser struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Photo    string `json:"photo" form:"photo"`
	Age      int    `json:"age" form:"age"`
	Bio      string `json:"bio" form:"bio"`
}

type AllUserFollow struct {
	ID       uint         `gorm:"primary_key"`
	Username string       `json:"username" form:"username"`
	Email    string       `json:"email" form:"email"`
	Photo    string       `json:"photo" form:"photo"`
	Age      int          `json:"age" form:"age"`
	Bio      string       `json:"bio" form:"bio"`
	Threads  []ThreadUser `json:"threads"`
}

func ConvertUserToAllUser(user *User) AllUser {
	return AllUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Photo:    user.Photo,
		Age:      user.Age,
		Bio:      user.Bio,
	}
}
func ConvertUserToAllUserFollow(user *User) AllUserFollow {
	threads := make([]ThreadUser, len(user.Threads))
	for i, thread := range user.Threads {
		threads[i] = ConvertThreadUser(&thread)
	}
	return AllUserFollow{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Photo:    user.Photo,
		Bio:      user.Bio,
		Age:      user.Age,
		Threads:  threads,
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
