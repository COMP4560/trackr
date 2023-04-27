package db

import (
	"fmt"
	"gorm.io/gorm"
	"trackr/src/models"
)

type OrganizationService struct {
	DB *gorm.DB
}

func (service *OrganizationService) GetUserOrganizations(user models.User) ([]models.UserOrganization, error) {
	var organizations []models.UserOrganization

	if result := service.DB.
		Preload("User").
		Preload("Organization").
		Find(&organizations, "user_organizations.user_id = ?", user.ID); result.Error != nil {
		return nil, result.Error
	}

	return organizations, nil
}

func (service *OrganizationService) GetUserOrganization(id uint, user models.User) (*models.UserOrganization, error) {
	var organizationWithRelation models.UserOrganization

	if result := service.DB.
		Preload("User").
		Preload("Organization").
		First(&organizationWithRelation, "user_organizations.user_id = ? AND user_organizations.organization_id = ?", user.ID, id); result.Error != nil {
		return nil, result.Error
	}

	return &organizationWithRelation, nil
}

func (service *OrganizationService) GetOrganizationByAPIKey(apiKey string) (*models.Organization, error) {
	var organization models.Organization

	if result := service.DB.
		Preload("Users").
		Table("user_organizations").
		Joins("INNER JOIN organizations ON user_organizations.organization_id = organization.id").
		First(&organization, "user_organizations.api_key = ?", apiKey); result.Error != nil {
		return nil, result.Error
	}

	return &organization, nil
}

func (service *OrganizationService) AddOrganization(organization models.Organization, userOrganization models.UserOrganization) (uint, error) {
	if result := service.DB.Create(&organization); result.Error != nil {
		return 0, result.Error
	}

	userOrganization.OrganizationID = organization.ID
	if result := service.DB.Create(&userOrganization); result.Error != nil {
		return 0, result.Error
	}

	return organization.ID, nil
}

func (service *OrganizationService) UpdateOrganization(organization models.Organization) error {
	if result := service.DB.Save(&organization); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *OrganizationService) DeleteOrganization(id uint, user models.User) error {
	if result := service.DB.
		Model(&models.Organization{}).
		Joins("INNER JOIN user_organizations ON user_organization.user_id = ? "+
			"AND user_organizations.role = ?", user.ID, "organization_owner").
		Delete(&models.Organization{}, "organizations.id = ?", id); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
