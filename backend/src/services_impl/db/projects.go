package db

import (
	"fmt"
	"gorm.io/gorm"
	"trackr/src/models"
)

type ProjectService struct {
	DB *gorm.DB
}

func (service *ProjectService) GetUserProjects(user models.User) ([]models.UserProject, error) {
	var userProjects []models.UserProject

	if result := service.DB.
		Preload("User").
		Preload("Project").
		Find(&userProjects, "user_projects.user_id = ?", user.ID); result.Error != nil {
		return nil, result.Error
	}

	return userProjects, nil
}

func (service *ProjectService) GetUserProject(id uint, user models.User) (*models.UserProject, error) {
	var userProject models.UserProject

	if result := service.DB.
		Preload("User").
		Preload("Project").
		First(&userProject, "user_projects.user_id = ? AND user_projects.project_id = ?", user.ID, id); result.Error != nil {
		return nil, result.Error
	}

	if userProject.ProjectID == 0 {
		return nil, fmt.Errorf("no project found")
	}

	return &userProject, nil
}

func (service *ProjectService) GetUserAndProjectByAPIKey(apiKey string) (*models.UserProject, error) {
	var userProject models.UserProject

	if result := service.DB.
		Preload("User").
		Preload("Project").
		Find(&userProject, "user_projects.api_key = ?", apiKey); result.Error != nil {
		return nil, result.Error
	}

	if userProject.ProjectID == 0 {
		return nil, fmt.Errorf("no project found")
	}

	return &userProject, nil
}

func (service *ProjectService) AddProject(project models.Project, userProject models.UserProject) (uint, error) {
	if result := service.DB.Create(&project); result.Error != nil {
		return 0, result.Error
	}

	userProject.ProjectID = project.ID
	if result := service.DB.Create(&userProject); result.Error != nil {
		return 0, result.Error
	}

	return project.ID, nil
}

func (service *ProjectService) UpdateProject(project models.Project, userProject models.UserProject) error {
	if result := service.DB.Save(&project); result.Error != nil {
		return result.Error
	}

	if result := service.DB.Save(&userProject); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *ProjectService) DeleteProject(id uint, user models.User) error {
	var userProject models.UserProject

	if result := service.DB.Find(&userProject, "user_projects.user_id = ? AND user_projects.project_id = ? AND user_projects.role = 'project_owner'", user.ID, id); result.Error != nil || userProject.UserID != user.ID || userProject.ProjectID != id {
		return fmt.Errorf("user does not have permission to remove project")
	}

	if result := service.DB.Delete(&models.Project{}, "id = ?", id); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	if result := service.DB.Delete(&models.UserProject{}, "project_id = ? AND user_id", id, user.ID); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	service.DB.Unscoped().Delete(&models.Field{}, "project_id = ?", id)

	return nil
}
