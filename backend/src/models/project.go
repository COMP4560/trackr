package models

import "time"

type Project struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	APIKey      string `gorm:"uniqueIndex"`
	UpdatedAt   time.Time
	CreatedAt   time.Time

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}
