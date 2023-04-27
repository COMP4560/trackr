package models

import (
	"database/sql"
	"time"
)

type UserOrganization struct {
	APIKey string `gorm:"uniqueIndex"`
	Role   string

	UserID uint `gorm:"primary_key"`
	User   User

	OrganizationID uint `gorm:"primary_key"`
	Organization   Organization

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
