package models

import (
	"time"
)

type Field struct {
	ID   uint `gorm:"primarykey"`
	Name string

	ProjectID uint
	Project   Project `gorm:"constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
