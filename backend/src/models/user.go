package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"uniqueIndex"`
	Password   string
	FirstName  string
	LastName   string
	IsVerified bool

	MaxValues        int64
	MaxValueInterval int64

	Organizations []*Organization `gorm:"many2many:user_organizations;"`
	Projects      []*Project      `gorm:"many2many:user_projects;"`
}

func (user *User) BeforeSave(db *gorm.DB) error {
	var organizationCount int64

	if result := db.
		Table("user_projects").
		Joins("INNER JOIN projects ON user_projects.project_id = projects.id").
		Joins("INNER JOIN user_organizations ON user_organizations.user_id = user_projects.user_id AND user_organizations.organization_id = projects.organization_id").
		Where("user_projects.user_id = ?", user.ID).
		Count(&organizationCount); result.Error != nil {
		return result.Error
	}

	if organizationCount != int64(len(user.Projects)) {
		return fmt.Errorf("user must be a member of the organization that the project belongs to")
	}

	return nil
}
