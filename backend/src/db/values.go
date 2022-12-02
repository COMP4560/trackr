package db

import (
	"fmt"
	"gorm.io/gorm"

	"trackr/src/models"
)

type ValueServiceDB struct {
	database *gorm.DB
}

func (service *ValueServiceDB) GetValues(field models.Field, user models.User, order string, offset int, limit int) ([]models.Value, error) {
	var values []models.Value

	result := service.database.Model(&models.Value{})
	result = result.Joins("LEFT JOIN fields ON `values`.`field_id` = `fields`.`id`")
	result = result.Joins("LEFT JOIN projects ON `fields`.`project_id` = `projects`.`id`")

	if order == "asc" {
		result = result.Order("`values`.`created_at` ASC")
	} else if order == "desc" {
		result = result.Order("`values`.`created_at` DESC")
	} else {
		return nil, fmt.Errorf("invalid order")
	}

	if offset > 0 {
		result = result.Offset(offset)
	}

	if limit > 0 {
		result = result.Limit(limit)
	}

	result = result.Find(&values, "`values`.`field_id` = ? AND `projects`.`user_id` = ?", field.ID, user.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return values, nil
}

func (service *ValueServiceDB) GetNumberOfValuesByUser(user models.User) (int64, error) {
	var count int64

	result := service.database.Model(&models.User{})
	result = result.Model(&models.Value{})
	result = result.Joins("LEFT JOIN fields ON `values`.`field_id` = `fields`.`id`")
	result = result.Joins("LEFT JOIN projects ON `fields`.`project_id` = `projects`.`id`")
	result = result.Joins("LEFT JOIN users ON `projects`.`user_id` = `users`.`id`")
	result = result.Where("`users`.`id` = ?", user.ID)
	result = result.Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *ValueServiceDB) GetNumberOfValuesByField(field models.Field) (int64, error) {
	var count int64

	result := service.database.Model(&models.User{})
	result = result.Model(&models.Value{})
	result = result.Joins("LEFT JOIN fields ON `values`.`field_id` = `fields`.`id`")
	result = result.Where("`fields`.`id` = ?", field.ID)
	result = result.Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *ValueServiceDB) GetLastAddedValue(user models.User) (*models.Value, error) {
	var value models.Value

	result := service.database.Model(&models.Value{})
	result = result.Joins("LEFT JOIN fields ON `values`.`field_id` = `fields`.`id`")
	result = result.Joins("LEFT JOIN projects ON `fields`.`project_id` = `projects`.`id`")
	result = result.Order("`values`.`created_at` DESC")
	result = result.First(&value, "`projects`.`user_id` = ?", user.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &value, nil
}

func (service *ValueServiceDB) GetValue(id uint, user models.User) (*models.Value, error) {
	var value models.Value

	result := service.database.Model(&models.Value{})
	result = result.Joins("LEFT JOIN fields ON `values`.`field_id` = `fields`.`id`")
	result = result.Joins("LEFT JOIN projects ON `fields`.`project_id` = `projects`.`id`")
	result = result.First(&value, "`values`.`id` = ? AND `projects`.`user_id` = ?", id, user.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &value, nil
}

func (service *ValueServiceDB) AddValue(value models.Value) error {
	if result := service.database.Create(&value); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *ValueServiceDB) DeleteValues(field models.Field) error {
	result := service.database.Delete(&models.Value{}, "field_id = ?", field.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
