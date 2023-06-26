package models

type ThreadLikeAssoc struct {
	UserID   uint `gorm:"primary_key"`
	ThreadID uint `gorm:"primary_key"`
}
