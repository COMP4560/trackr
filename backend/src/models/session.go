package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	ID        string `gorm:"primarykey"`
	UserAgent string
	ExpiresAt time.Time

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}
