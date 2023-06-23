package models

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	UserID   int    `json:"user_id" form:"user_id"`
	ThreadID int    `json:"thread_id" form:"thread_id" validate:"required"`
	Thread   Thread `json:"thread"`
}

type Followrespon struct {
	ThreadID int          `json:"thread_id" form:"thread_id"`
	UserID   int          `json:"user_id" form:"user_id"`
	Thread   []ThreadUser `json:"thread"`
}

// func ConvertFollow(Follow *Follow) Followrespon {
// 	Follows := make([]ThreadUser, len(Follow.Thread))
// 	for i, follow := range Follow.Thread {
// 		Follows[i] = ConvertThreadUser(&follow)
// 	}

// 	return Followrespon{
// 		ThreadID: Follow.ThreadID,
// 		UserID:   Follow.UserID,
// 		Thread:   Follows,
// 	}
// }
