package db

import (
	"time"

	"gorm.io/gorm"

	"trackr/src/models"
)

type LogServiceDB struct {
	database *gorm.DB
}

func (service *LogServiceDB) GetLogs(user models.User) ([]models.Log, error) {
	var logs []models.Log
	if result := service.database.Preload("Project").Order("created_at DESC").Find(&logs, "user_id = ?", user.ID); result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func (service *LogServiceDB) AddLog(message string, user models.User, projectId *uint) error {
	log := models.Log{
		Message:   message,
		CreatedAt: time.Now(),

		User:      user,
		ProjectID: projectId,
	}

	if result := service.database.Create(&log); result.Error != nil {
		return result.Error
	}

	return nil
}
