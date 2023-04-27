package services

import (
	"trackr/src/models"
)

type ProjectService interface {
	GetUserProjects(user models.User) ([]models.UserProject, error)
	GetUserProject(id uint, user models.User) (*models.UserProject, error)
	GetUserAndProjectByAPIKey(apiKey string) (*models.UserProject, error)
	AddProject(project models.Project, userProject models.UserProject) (uint, error)
	UpdateProject(project models.Project, userProject models.UserProject) error
	DeleteProject(id uint, user models.User) error
}
