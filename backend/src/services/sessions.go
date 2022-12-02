package services

import "trackr/src/models"

type SessionService interface {
	GetSessionAndUser(id string) (*models.Session, *models.User, error)
	AddSession(session models.Session) error
	DeleteSession(id string, user models.User) error
	DeleteExpiredSessions(user models.User) error
}
