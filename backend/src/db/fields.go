package db

import (
	"fmt"
	"gorm.io/gorm"

	"trackr/src/models"
)

type FieldServiceDB struct {
	database *gorm.DB
}

func (service *FieldServiceDB) GetFields(project models.Project, user models.User) ([]models.Field, error) {
	var fields []models.Field
	if result := service.database.Model(&models.Field{}).Joins(
		"LEFT JOIN projects ON fields.project_id = projects.id",
	).Find(
		&fields, "fields.project_id = ? AND projects.user_id = ?", project.ID, user.ID,
	); result.Error != nil {
		return nil, result.Error
	}

	return fields, nil
}

func (service *FieldServiceDB) GetNumberOfFieldsByProject(project models.Project, user models.User) (int64, error) {
	var count int64

	if result := service.database.Model(&models.Field{}).Joins(
		"LEFT JOIN projects ON fields.project_id = projects.id",
	).Where("fields.project_id = ? AND projects.user_id = ?", project.ID, user.ID).Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *FieldServiceDB) GetNumberOfFieldsByUser(user models.User) (int64, error) {
	var count int64

	if result := service.database.Model(&models.Field{}).Joins(
		"LEFT JOIN projects ON fields.project_id = projects.id",
	).Where("projects.user_id = ?", user.ID).Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *FieldServiceDB) GetField(id uint, user models.User) (*models.Field, error) {
	var field models.Field
	if result := service.database.Model(&models.Field{}).Joins(
		"LEFT JOIN projects ON fields.project_id = projects.id",
	).First(
		&field, "fields.id = ? AND projects.user_id = ?", id, user.ID,
	); result.Error != nil {
		return nil, result.Error
	}

	return &field, nil
}

func (service *FieldServiceDB) AddField(field models.Field) (uint, error) {
	if result := service.database.Create(&field); result.Error != nil {
		return 0, result.Error
	}

	return field.ID, nil
}

func (service *FieldServiceDB) UpdateField(field models.Field) error {
	if result := service.database.Save(&field); result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *FieldServiceDB) DeleteField(field models.Field) error {
	result := service.database.Delete(field)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
