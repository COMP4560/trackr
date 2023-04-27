package db

import (
	"fmt"
	"gorm.io/gorm"
	"trackr/src/models"
)

type FieldService struct {
	DB *gorm.DB
}

func (service *FieldService) GetFields(project models.Project, user models.User) ([]models.Field, error) {
	var fields []models.Field

	if result := service.DB.
		Model(&models.Field{}).
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id").
		Find(&fields, "user_projects.project_id = ? AND user_projects.user_id = ?", project.ID, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return fields, nil
}

func (service *FieldService) GetNumberOfProjectsByOrganization(organization models.Organization, user models.User) (int64, error) {
	var count int64

	if result := service.DB.
		Model(&models.User{}).
		Joins("INNER JOIN user_projects ON user_projects.user_id = users.id").
		Joins("INNER JOIN projects ON user_projects.project_id = projects.id").
		Where("projects.organization_id = ? AND users.id = ?", organization.ID, user.ID).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *FieldService) GetNumberOfFieldsByProject(project models.Project, user models.User) (int64, error) {
	var count int64

	if result := service.DB.
		Model(&models.Field{}).
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id AND user_projects.user_id").
		Where("user_projects.project_id = ? AND user_projects.user_id = ?", project.ID, user.ID).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *FieldService) GetNumberOfFieldsByUser(user models.User) (int64, error) {
	var count int64

	if result := service.DB.
		Model(&models.Field{}).
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id AND user_projects.user_id").
		Where("user_projects.user_id = ?", user.ID).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *FieldService) GetField(id uint, user models.User) (*models.Field, error) {
	var field models.Field

	if result := service.DB.
		Model(&models.Field{}).
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id AND user_projects.user_id").
		Where("user_projects.user_id = ?", user.ID).
		First(&field, "fields.id = ?", id); result.Error != nil {
		return nil, result.Error
	}

	return &field, nil
}

func (service *FieldService) AddField(field models.Field) (uint, error) {
	if result := service.DB.Create(&field); result.Error != nil {
		return 0, result.Error
	}

	return field.ID, nil
}

func (service *FieldService) UpdateField(field models.Field) error {
	if result := service.DB.Save(&field); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *FieldService) DeleteField(field models.Field) error {
	if result := service.DB.Delete(&models.Field{}, "id = ?", field.ID); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	if result := service.DB.Where("field_id = ?", field.ID).Delete(&models.Value{}); result.Error != nil {
		return result.Error
	}

	if result := service.DB.Where("field_id = ?", field.ID).Delete(&models.Visualization{}); result.Error != nil {
		return result.Error
	}

	return nil
}
