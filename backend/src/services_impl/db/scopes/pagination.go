package scopes

import (
	"gorm.io/gorm"
)

type OrganizationService struct {
	DB *gorm.DB
}

func SetPagination(offset int, limit int) func(database *gorm.DB) *gorm.DB {
	return func(database *gorm.DB) *gorm.DB {

		if offset > 0 {
			database = database.Offset(offset)
		}

		if limit > 0 {
			database = database.Limit(limit)
		}

		return database
	}
}
