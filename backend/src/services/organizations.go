package services

import (
	"trackr/src/models"
)

type OrganizationService interface {
	GetUserOrganizations(user models.User) ([]models.UserOrganization, error)
	GetUserOrganization(id uint, user models.User) (*models.UserOrganization, error)
	GetOrganizationByAPIKey(apiKey string) (*models.Organization, error)
	AddOrganization(organization models.Organization, userOrganization models.UserOrganization) (uint, error)
	UpdateOrganization(organization models.Organization) error
	DeleteOrganization(id uint, user models.User) error
}
