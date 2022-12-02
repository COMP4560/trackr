package db

import (
	"fmt"
	"gorm.io/gorm"

	"trackr/src/models"
)

type UserServiceDB struct {
	database *gorm.DB
}

func (service *UserServiceDB) GetUser(email string) (*models.User, error) {
	var user models.User
	if result := service.database.First(&user, "email = ?", email); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (service *UserServiceDB) GetNumberOfUsers(email string) (int64, error) {
	var count int64
	if result := service.database.Model(&models.User{}).Where("email = ?", email).Count(&count); result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (service *UserServiceDB) AddUser(user models.User) (uint, error) {
	if result := service.database.Create(&user); result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (service *UserServiceDB) DeleteUser(user models.User) error {
	result := service.database.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (service *UserServiceDB) UpdateUser(user models.User) error {
	if result := service.database.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
