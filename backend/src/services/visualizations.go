package services

import "trackr/src/models"

type VisualizationService interface {
	GetVisualizations(project models.Project, user models.User) ([]models.Visualization, error)
	GetVisualization(id uint, user models.User) (*models.Visualization, error)
	AddVisualization(visualization models.Visualization) (uint, error)
	UpdateVisualization(visualization models.Visualization) error
	DeleteVisualization(visualization models.Visualization) error
}
