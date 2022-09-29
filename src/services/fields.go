package services

import "trackr/src/models"

type FieldService interface {
	GetFields(project models.Project, user models.User) ([]models.Field, error)
	GetField(id uint, user models.User) (*models.Field, error)
	AddField(field models.Field) error
	UpdateField(field models.Field) error
	DeleteField(id uint, user models.User) error
}
