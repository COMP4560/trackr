package db

import (
	"gorm.io/gorm"

	"trackr/src/models"
)

type FieldServiceDB struct {
	database *gorm.DB
}

func (service *FieldServiceDB) GetFields(project models.Project, user models.User) ([]models.Field, error) {
	var fields []models.Field
	if result := service.database.Model(&models.Field{}).Joins("LEFT JOIN projects").Find(&fields, "fields.project_id = ? AND projects.user_id = ?", project.ID, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return fields, nil
}

func (service *FieldServiceDB) GetField(id uint, user models.User) (*models.Field, error) {
	var field models.Field
	if result := service.database.Model(&models.Field{}).Joins("LEFT JOIN projects").First(&field, "fields.id = ? AND projects.user_id = ?", id, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return &field, nil
}

func (service *FieldServiceDB) AddField(field models.Field) error {
	if result := service.database.Create(&field); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *FieldServiceDB) UpdateField(field models.Field) error {
	if result := service.database.Save(&field); result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *FieldServiceDB) DeleteField(id uint, user models.User) error {
	field, err := service.GetField(id, user)
	if err != nil {
		return err
	}

	if result := service.database.Delete(field); result.Error != nil {
		return result.Error
	}

	return nil
}
