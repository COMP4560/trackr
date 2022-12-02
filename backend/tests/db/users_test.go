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
