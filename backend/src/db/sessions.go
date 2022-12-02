package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"trackr/src/models"
)

type SessionServiceDB struct {
	database *gorm.DB
}

func (service *SessionServiceDB) GetSessionAndUser(id string) (*models.Session, *models.User, error) {
	var session models.Session
	if result := service.database.Preload("User").First(&session, "id = ?", id); result.Error != nil {
		return nil, nil, result.Error
	}

	return &session, &session.User, nil
}

func (service *SessionServiceDB) AddSession(session models.Session) error {
	if result := service.database.Create(&session); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *SessionServiceDB) DeleteSession(id string, user models.User) error {
	result := service.database.Delete(&models.Session{}, "id = ? AND user_id = ?", id, user.ID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (service *SessionServiceDB) DeleteExpiredSessions(user models.User) error {
	result := service.database.Delete(&models.Session{}, "expires_at < ? AND user_id = ?", time.Now(), user.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
