package db

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"trackr/src/models"
)

type SessionService struct {
	DB *gorm.DB
}

func (service *SessionService) GetSessionAndUser(id string) (*models.Session, *models.User, error) {
	var session models.Session

	if result := service.DB.
		Preload("User").
		First(&session, "id = ?", id); result.Error != nil {
		return nil, nil, result.Error
	}

	return &session, &session.User, nil
}

func (service *SessionService) AddSession(session models.Session) error {
	if result := service.DB.Create(&session); result.Error != nil {
		return result.Error
	}

	return nil
}

func (service *SessionService) DeleteSession(id string, user models.User) error {
	result := service.DB.Delete(&models.Session{}, "id = ? AND user_id = ?", id, user.ID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (service *SessionService) DeleteExpiredSessions(user models.User) error {
	if result := service.DB.Delete(&models.Session{}, "expires_at < ? AND user_id = ?", time.Now(), user.ID); result.Error != nil {
		return result.Error
	}

	return nil
}
