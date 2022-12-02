package services

import (
	"trackr/src/models"
)

type LogService interface {
	GetLogs(user models.User) ([]models.Log, error)
	AddLog(message string, user models.User, projectId *uint) error
}
