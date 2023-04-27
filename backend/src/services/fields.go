package services

import (
	"trackr/src/models"
)

type FieldService interface {
	GetFields(project models.Project, user models.User) ([]models.Field, error)
	GetNumberOfProjectsByOrganization(project models.Organization, user models.User) (int64, error)
	GetNumberOfFieldsByProject(project models.Project, user models.User) (int64, error)
	GetNumberOfFieldsByUser(user models.User) (int64, error)
	GetField(id uint, user models.User) (*models.Field, error)
	AddField(field models.Field) (uint, error)
	UpdateField(field models.Field) error
	DeleteField(field models.Field) error
}
