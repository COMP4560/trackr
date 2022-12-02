package services

import "trackr/src/models"

type ProjectService interface {
	GetProjects(user models.User) ([]models.Project, error)
	GetProject(id uint, user models.User) (*models.Project, error)
	GetProjectByAPIKey(apiKey string) (*models.Project, error)
	AddProject(project models.Project) (uint, error)
	UpdateProject(project models.Project) error
	DeleteProject(id uint, user models.User) error
}
