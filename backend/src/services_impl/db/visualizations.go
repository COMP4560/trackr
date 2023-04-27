package db

import (
	"fmt"
	"gorm.io/gorm"
	"trackr/src/models"
)

type VisualizationService struct {
	DB *gorm.DB
}

func (service *VisualizationService) GetVisualizations(project models.Project, user models.User) ([]models.Visualization, error) {
	var visualizations []models.Visualization

	if result := service.DB.
		Model(&models.Visualization{}).
		Preload("Field").
		Joins("INNER JOIN fields ON visualizations.field_id = fields.id").
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id AND user_projects.user_id").
		Find(&visualizations, "user_projects.project_id = ? AND user_projects.user_id = ?", project.ID, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return visualizations, nil
}

func (service *VisualizationService) GetVisualization(id uint, user models.User) (*models.Visualization, error) {
	var visualization models.Visualization

	if result := service.DB.
		Model(&models.Visualization{}).
		Joins("INNER JOIN fields ON visualizations.field_id = fields.id").
		Joins("INNER JOIN projects ON fields.project_id = projects.id").
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id AND user_projects.user_id").
		First(&visualization, "visualizations.id = ? AND user_projects.user_id = ?", id, user.ID); result.Error != nil {
		return nil, result.Error
	}

	return &visualization, nil
}

func (service *VisualizationService) AddVisualization(visualization models.Visualization) (uint, error) {
	if result := service.DB.Create(&visualization); result.Error != nil {
		return 0, result.Error
	}

	return visualization.ID, nil
}

func (service *VisualizationService) UpdateVisualization(visualization models.Visualization) error {
	if result := service.DB.Save(&visualization); result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *VisualizationService) DeleteVisualization(visualization models.Visualization) error {
	if result := service.DB.Delete(&visualization); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
