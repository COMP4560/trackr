package db

import (
	"fmt"
	"gorm.io/gorm"
	"trackr/src/models"
)

type UserService struct {
	DB *gorm.DB
}

func (service *UserService) GetUser(email string) (*models.User, error) {
	var user models.User

	if result := service.DB.First(&user, "email = ?", email); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (service *UserService) GetNumberOfUsers(email string) (int64, error) {
	var count int64

	if result := service.DB.
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *UserService) AddUser(user models.User) (uint, error) {
	if result := service.DB.Create(&user); result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (service *UserService) DeleteUser(user models.User) error {
	if result := service.DB.Where("id = ?", user.ID).Delete(&models.User{}); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	if result := service.DB.
		Model(&models.Project{}).
		Joins("INNER JOIN user_projects ON user_projects.project_id = projects.id AND user_projects.user_id = ? AND user_projects.role = 'project_owner'", user.ID).
		Not("EXISTS (SELECT 1 FROM user_projects WHERE user_projects.project_id = user_projects.project_id AND user_id != ?)", user.ID).
		Delete(&models.Project{}); result.Error != nil {
		return result.Error
	}

	service.DB.Delete(&models.UserProject{}, "user_id = ?", user.ID) // TODO see if needed
	service.DB.Delete(&models.Session{}, "user_id = ?", user.ID)

	return nil
}

func (service *UserService) UpdateUser(user models.User) error {
	if result := service.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
