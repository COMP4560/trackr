package models

import "time"

type Log struct {
	ID        uint `gorm:"primarykey"`
	Message   string
	CreatedAt time.Time

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	ProjectID *uint
	Project   Project `gorm:"constraint:OnDelete:CASCADE;"`
}
