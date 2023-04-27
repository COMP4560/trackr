package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trackr/src/models"
	"trackr/tests"
)

func TestGetUser(t *testing.T) {
	suite := tests.Startup()

	user, err := suite.Service.GetUserService().GetUser("invalid@email")
	assert.NotNil(t, err)
	assert.Nil(t, user)

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, user.ID, suite.User.ID)
}

func TestGetNumberOfUsers(t *testing.T) {
	suite := tests.Startup()

	numberOfUsers, err := suite.Service.GetUserService().GetNumberOfUsers("invalid@email")
	assert.Nil(t, err)
	assert.Equal(t, numberOfUsers, int64(0))

	numberOfUsers, err = suite.Service.GetUserService().GetNumberOfUsers(suite.User.Email)
	assert.Nil(t, err)
	assert.Equal(t, numberOfUsers, int64(1))
}

func TestAddUser(t *testing.T) {
	suite := tests.Startup()

	numberOfUsers, err := suite.Service.GetUserService().GetNumberOfUsers("new@email")
	assert.Nil(t, err)
	assert.Equal(t, numberOfUsers, int64(0))

	user, err := suite.Service.GetUserService().GetUser("new@email")
	assert.NotNil(t, err)
	assert.Nil(t, user)

	_, err = suite.Service.GetUserService().AddUser(suite.User)
	assert.NotNil(t, err)

	newUser := suite.User
	newUser.ID = 2
	newUser.Email = "new@email"

	newUserID, err := suite.Service.GetUserService().AddUser(newUser)
	assert.Nil(t, err)

	numberOfUsers, err = suite.Service.GetUserService().GetNumberOfUsers(newUser.Email)
	assert.Nil(t, err)
	assert.Equal(t, numberOfUsers, int64(1))

	user, err = suite.Service.GetUserService().GetUser(newUser.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, newUser.ID, newUserID)
	assert.Equal(t, newUser.ID, user.ID)
}

func TestDeleteUser(t *testing.T) {
	suite := tests.Startup()

	//
	// Check before we delete if our relationships are correct.
	//

	userProject, err := suite.Service.GetProjectService().GetUserProject(suite.Project.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, userProject)
	assert.NotNil(t, userProject.Project)
	assert.NotNil(t, userProject.User)

	session, _, err := suite.Service.GetSessionService().GetSessionAndUser(suite.Session.ID)
	assert.Nil(t, err)
	assert.NotNil(t, session)

	field, err := suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, field)

	value, err := suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)

	visualization, err := suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs), len(logs))

	//
	// Begin delete test.
	//

	numberOfUsers, err := suite.Service.GetUserService().GetNumberOfUsers(suite.User.Email)
	assert.Nil(t, err)
	assert.Equal(t, numberOfUsers, int64(1))

	err = suite.Service.GetUserService().DeleteUser(models.User{})
	assert.NotNil(t, err)

	err = suite.Service.GetUserService().DeleteUser(suite.User)
	assert.Nil(t, err)

	numberOfUsers, err = suite.Service.GetUserService().GetNumberOfUsers(suite.User.Email)
	assert.Nil(t, err)
	assert.Equal(t, numberOfUsers, int64(0))

	user, err := suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.NotNil(t, err)
	assert.Nil(t, user)

	err = suite.Service.GetUserService().DeleteUser(suite.User)
	assert.NotNil(t, err)

	//
	// Check if cascade delete works.
	//

	userProject, err = suite.Service.GetProjectService().GetUserProject(suite.Project.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, userProject)

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser(suite.Session.ID)
	assert.NotNil(t, err)
	assert.Nil(t, session)

	field, err = suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, field)

	value, err = suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, visualization)
}

func TestUpdateUser(t *testing.T) {
	suite := tests.Startup()

	user, err := suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, user.FirstName, suite.User.FirstName)

	newUser := suite.User
	newUser.FirstName = "FirstName2"

	err = suite.Service.GetUserService().UpdateUser(newUser)
	assert.Nil(t, err)

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, user.FirstName, newUser.FirstName)
}
