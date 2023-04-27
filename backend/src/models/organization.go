package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string
	Description string

	Users    []*User    `gorm:"many2many:user_organizations;"`
	Projects []*Project `gorm:"foreignKey:OrganizationID"`
}
