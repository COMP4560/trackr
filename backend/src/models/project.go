package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string
	Description string

	Users          []*User       `gorm:"many2many:user_projects;"`
	OrganizationID sql.NullInt64 // uint
}
