package db

import (
	"fmt"
	"gorm.io/gorm"

	"trackr/src/models"
)

type VisulizationServiceDB struct {
	database *gorm.DB
}

func (service *VisulizationServiceDB) GetVisualizations(project models.Project, user models.User) ([]models.Visualization, error) {
	var visualizations []models.Visualization

	result := service.database.Model(&models.Visualization{})
	result = result.Preload("Field")
	result = result.Joins("LEFT JOIN fields ON `visualizations`.`field_id` = `fields`.`id`")
	result = result.Joins("LEFT JOIN projects ON `fields`.`project_id` = `projects`.`id`")
	result = result.Find(&visualizations, "`projects`.`id` = ? AND `projects`.`user_id` = ?", project.ID, user.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return visualizations, nil
}

func (service *VisulizationServiceDB) GetVisualization(id uint, user models.User) (*models.Visualization, error) {
	var visualization models.Visualization

	result := service.database.Model(&models.Visualization{})
	result = result.Joins("LEFT JOIN fields ON `visualizations`.`field_id` = `fields`.`id`")
	result = result.Joins("LEFT JOIN projects ON `fields`.`project_id` = `projects`.`id`")
	result = result.First(&visualization, "`visualizations`.`id` = ? AND `projects`.`user_id` = ?", id, user.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &visualization, nil
}

func (service *VisulizationServiceDB) AddVisualization(visualization models.Visualization) (uint, error) {
	if result := service.database.Create(&visualization); result.Error != nil {
		return 0, result.Error
	}

	return visualization.ID, nil
}

func (service *VisulizationServiceDB) UpdateVisualization(visualization models.Visualization) error {
	if result := service.database.Save(&visualization); result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *VisulizationServiceDB) DeleteVisualization(visualization models.Visualization) error {
	result := service.database.Delete(visualization)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
