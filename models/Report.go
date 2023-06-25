package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Reason       string
	ThreadID     uint
	ReportThread Thread `gorm:"foreignKey:ThreadID;references:ID"`
	UserID       uint
	ReportUser   User `gorm:"foreignKey:UserID;references:ID"`
}
