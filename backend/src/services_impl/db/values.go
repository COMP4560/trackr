package db

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"trackr/src/models"
	"trackr/src/services_impl/db/scopes"
)

type ValueService struct {
	DB *gorm.DB
}

func (service *ValueService) GetValues(field models.Field, user models.User, order string, offset int, limit int) ([]models.Value, error) {
	var values []models.Value
	order = strings.ToUpper(order)

	if order != "ASC" && order != "DESC" {
		return nil, fmt.Errorf("invalid order")
	}

	if result := service.DB.
		Model(&models.Value{}).
		Joins("INNER JOIN fields ON `values`.field_id = fields.id").
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id").
		Order("`values`.created_at "+order).
		Scopes(scopes.SetPagination(offset, limit)).
		Find(&values, "`values`.field_id = ? AND user_projects.user_id = ?", field.ID, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return values, nil
}

func (service *ValueService) GetNumberOfValuesByUser(user models.User) (int64, error) {
	var count int64

	if result := service.DB.
		Model(&models.Value{}).
		Joins("INNER JOIN fields ON `values`.field_id = fields.id").
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id").
		Where("user_projects.user_id = ?", user.ID).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *ValueService) GetNumberOfValuesByField(field models.Field) (int64, error) {
	var count int64

	if result := service.DB.
		Model(&models.Value{}).
		Joins("INNER JOIN fields ON `values`.field_id = fields.id").
		Where("fields.id = ?", field.ID).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *ValueService) GetLastAddedValue(user models.User) (*models.Value, error) {
	var value models.Value

	if result := service.DB.
		Model(&models.Value{}).
		Joins("INNER JOIN fields ON `values`.field_id = fields.id").
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id").
		Order("`values`.created_at DESC").
		First(&value, "user_projects.user_id = ?", user.ID); result.Error != nil {
		return nil, result.Error
	}

	return &value, nil
}

func (service *ValueService) GetValue(id uint, user models.User) (*models.Value, error) {
	var value models.Value

	var values []models.Value
	service.DB.Find(&values)

	var fields []models.Field
	service.DB.Find(&fields)

	if result := service.DB.
		Model(&models.Value{}).
		Joins("INNER JOIN fields ON `values`.field_id = fields.id").
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id").
		First(&value, "`values`.id = ? AND user_projects.user_id = ?", id, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return &value, nil
}

func (service *ValueService) AddValue(value models.Value) error {
	if result := service.DB.Create(&value); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *ValueService) DeleteValues(field models.Field) error {
	if result := service.DB.Delete(&models.Value{}, "field_id = ?", field.ID); result.Error != nil {
		return result.Error
	}

	return nil
}
