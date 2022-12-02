package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"trackr/src/models"
	"trackr/tests"
)

func TestGetLogs(t *testing.T) {
	suite := tests.Startup()

	logs, err := suite.Service.GetLogService().GetLogs(models.User{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(logs))

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs), len(logs))

	for i, log := range suite.Logs {
		assert.Equal(t, log.ID, logs[i].ID)
		assert.Equal(t, log.Message, logs[i].Message)
		assert.Equal(t, log.ProjectID, logs[i].ProjectID)

		if logs[i].ProjectID != nil {
			assert.Equal(t, log.Project.ID, logs[i].Project.ID)
			assert.Equal(t, log.Project.Name, logs[i].Project.Name)
		}
	}
}

func TestAddLog(t *testing.T) {
	suite := tests.Startup()

	err := suite.Service.GetLogService().AddLog("Invalid Log", models.User{}, nil)
	assert.NotNil(t, err)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs), len(logs))

	projectId := models.Project{}.ID
	err = suite.Service.GetLogService().AddLog("Invalid Log", suite.User, &projectId)
	assert.NotNil(t, err)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs), len(logs))

	err = suite.Service.GetLogService().AddLog("Third Log", suite.User, nil)
	assert.Nil(t, err)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs)+1, len(logs))
	assert.Equal(t, "Third Log", logs[0].Message)
	assert.Nil(t, logs[0].ProjectID)

	err = suite.Service.GetLogService().AddLog("Fourth Log with Project", suite.User, &suite.Project.ID)
	assert.Nil(t, err)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs)+2, len(logs))
	assert.Equal(t, "Fourth Log with Project", logs[0].Message)
	assert.NotNil(t, logs[0].ProjectID)
	assert.Equal(t, suite.Project.ID, logs[0].Project.ID)
	assert.Equal(t, suite.Project.Name, logs[0].Project.Name)
}
