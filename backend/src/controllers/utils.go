package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
	"trackr/src/models"

	"trackr/src/common"
)

func generateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func generateAPIKey() (string, error) {
	return common.RandomString(apiKeyLength)
}

func getLoggedInUser(c *gin.Context) *models.User {
	return c.MustGet("user").(*models.User)
}

func isLoggedIn(c *gin.Context) *models.User {
	sessionId, err := c.Cookie(sessionCookie)
	if err != nil {
		return nil
	}

	session, user, err := serviceProvider.GetSessionService().GetSessionAndUser(sessionId)
	if err != nil {
		return nil
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil
	}

	return user
}
