package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Message string

	UserID uint
	User   User

	ProjectID *uint
	Project   Project `gorm:"constraint:OnDelete:SET NULL;"`
}
