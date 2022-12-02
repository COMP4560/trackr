package services

import "trackr/src/models"

type UserService interface {
	GetUser(email string) (*models.User, error)
	GetNumberOfUsers(email string) (int64, error)
	AddUser(user models.User) (uint, error)
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}
