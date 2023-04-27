package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trackr/src/models"
	"trackr/tests"
)

func TestGetSessionAndUser(t *testing.T) {
	suite := tests.Startup()

	session, user, err := suite.Service.GetSessionService().GetSessionAndUser("InvalidSessionID")
	assert.NotNil(t, err)
	assert.Nil(t, session)
	assert.Nil(t, user)

	session, user, err = suite.Service.GetSessionService().GetSessionAndUser(suite.Session.ID)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.NotNil(t, user)

	assert.Equal(t, user.ID, suite.User.ID)
	assert.Equal(t, session.ID, suite.Session.ID)
}

func TestAddSession(t *testing.T) {
	suite := tests.Startup()

	err := suite.Service.GetSessionService().AddSession(suite.Session)
	assert.NotNil(t, err)

	newSession := suite.Session
	newSession.ID = "NewSessionID"

	session, user, err := suite.Service.GetSessionService().GetSessionAndUser(newSession.ID)
	assert.NotNil(t, err)
	assert.Nil(t, session)
	assert.Nil(t, user)

	err = suite.Service.GetSessionService().AddSession(newSession)
	assert.Nil(t, err)

	session, user, err = suite.Service.GetSessionService().GetSessionAndUser(newSession.ID)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.NotNil(t, user)

	assert.Equal(t, user.ID, newSession.User.ID)
	assert.Equal(t, session.ID, newSession.ID)
}

func TestDeleteSession(t *testing.T) {
	suite := tests.Startup()

	session, _, err := suite.Service.GetSessionService().GetSessionAndUser(suite.Session.ID)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, session.ID, suite.Session.ID)

	err = suite.Service.GetSessionService().DeleteSession("InvalidSessionID", models.User{})
	assert.NotNil(t, err)

	err = suite.Service.GetSessionService().DeleteSession(suite.Session.ID, models.User{})
	assert.NotNil(t, err)

	err = suite.Service.GetSessionService().DeleteSession("InvalidSessionID", suite.User)
	assert.NotNil(t, err)

	err = suite.Service.GetSessionService().DeleteSession(suite.Session.ID, suite.User)
	assert.Nil(t, err)

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser(suite.Session.ID)
	assert.NotNil(t, err)
	assert.Nil(t, session)
}

func TestDeleteExpiredSessions(t *testing.T) {
	suite := tests.Startup()

	session, _, err := suite.Service.GetSessionService().GetSessionAndUser(suite.ExpiredSession.ID)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, session.ID, suite.ExpiredSession.ID)

	err = suite.Service.GetSessionService().DeleteExpiredSessions(models.User{})
	assert.Nil(t, err)

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser(suite.ExpiredSession.ID)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, session.ID, suite.ExpiredSession.ID)

	err = suite.Service.GetSessionService().DeleteExpiredSessions(suite.User)
	assert.Nil(t, err)

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser(suite.ExpiredSession.ID)
	assert.NotNil(t, err)
	assert.Nil(t, session)
}
