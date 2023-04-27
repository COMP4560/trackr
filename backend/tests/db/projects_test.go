package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trackr/src/models"
	"trackr/tests"
	"trackr/tests/helpers"
)

func TestGetProjects(t *testing.T) {
	suite := tests.Startup()

	projects, err := suite.Service.GetProjectService().GetUserProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 1)
	assert.Equal(t, projects[0].Project.ID, suite.Project.ID)

	projects, err = suite.Service.GetProjectService().GetUserProjects(models.User{})
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 0)
}

func TestGetUserProject(t *testing.T) {
	suite := tests.Startup()

	userProject, err := suite.Service.GetProjectService().GetUserProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, userProject)
	assert.NotNil(t, userProject.Project)
	assert.NotNil(t, userProject.User)

	assert.Equal(t, userProject.ProjectID, suite.Project.ID)
	assert.Equal(t, userProject.UserID, suite.User.ID)

	userProject, err = suite.Service.GetProjectService().GetUserProject(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, userProject)

	userProject, err = suite.Service.GetProjectService().GetUserProject(1, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, userProject)
}

func TestGetProjectAndUserByAPIKey(t *testing.T) {
	suite := tests.Startup()

	userProject, err := suite.Service.GetProjectService().GetUserAndProjectByAPIKey(suite.UserProject.APIKey)
	assert.Nil(t, err)
	assert.NotNil(t, userProject)
	assert.NotNil(t, userProject.Project)
	assert.NotNil(t, userProject.User)

	assert.Equal(t, userProject.ProjectID, suite.Project.ID)
	assert.Equal(t, userProject.UserID, suite.User.ID)

	userProject, err = suite.Service.GetProjectService().GetUserAndProjectByAPIKey(models.UserProject{}.APIKey)
	assert.NotNil(t, err)
	assert.Nil(t, userProject)
}

func TestAddProject(t *testing.T) {
	suite := tests.Startup()

	newProject, newUserProject := helpers.CreateNewProject(*suite)

	projectId, err := suite.Service.GetProjectService().AddProject(newProject, newUserProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	projects, err := suite.Service.GetProjectService().GetUserProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 2)
	assert.Equal(t, projects[1].Project.ID, newProject.ID)
}

func TestUpdateProject(t *testing.T) {
	suite := tests.Startup()

	newProject := suite.Project

	newUserProject := suite.UserProject
	newUserProject.APIKey = "APIKey2"

	err := suite.Service.GetProjectService().UpdateProject(suite.Project, newUserProject)
	assert.Nil(t, err)

	projects, err := suite.Service.GetProjectService().GetUserProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 1)
	assert.Equal(t, projects[0].Project.ID, newProject.ID)
	assert.Equal(t, projects[0].APIKey, newUserProject.APIKey)
}

func TestDeleteProject(t *testing.T) {
	suite := tests.Startup()

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
	assert.NotNil(t, logs[0].ProjectID)

	err = suite.Service.GetProjectService().DeleteProject(suite.Project.ID, models.User{})
	assert.NotNil(t, err)

	err = suite.Service.GetProjectService().DeleteProject(2, suite.User)
	assert.NotNil(t, err)

	err = suite.Service.GetProjectService().DeleteProject(suite.Project.ID, suite.User)
	assert.Nil(t, err)

	projects, err := suite.Service.GetProjectService().GetUserProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(projects))

	err = suite.Service.GetProjectService().DeleteProject(suite.Project.ID, suite.User)
	assert.NotNil(t, err)

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
