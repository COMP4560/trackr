package models

import (
	"database/sql"
	"time"
)

type UserProject struct {
	APIKey string `gorm:"uniqueIndex"`
	Role   string

	UserID uint `gorm:"primary_key"`
	User   User

	ProjectID uint `gorm:"primary_key"`
	Project   Project

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
