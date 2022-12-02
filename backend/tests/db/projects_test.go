package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"trackr/src/models"
	"trackr/tests"
)

func TestGetProjects(t *testing.T) {
	suite := tests.Startup()

	projects, err := suite.Service.GetProjectService().GetProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 1)
	assert.Equal(t, projects[0].ID, suite.Project.ID)

	projects, err = suite.Service.GetProjectService().GetProjects(models.User{})
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 0)
}

func TestGetProject(t *testing.T) {
	suite := tests.Startup()

	project, err := suite.Service.GetProjectService().GetProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, project.ID, suite.Project.ID)

	project, err = suite.Service.GetProjectService().GetProject(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, project)

	project, err = suite.Service.GetProjectService().GetProject(1, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, project)
}

func TestGetProjectByAPIKey(t *testing.T) {
	suite := tests.Startup()

	project, err := suite.Service.GetProjectService().GetProjectByAPIKey(suite.Project.APIKey)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, project.ID, suite.Project.ID)

	project, err = suite.Service.GetProjectService().GetProjectByAPIKey(models.Project{}.APIKey)
	assert.NotNil(t, err)
	assert.Nil(t, project)
}

func TestAddProject(t *testing.T) {
	suite := tests.Startup()

	newProject := suite.Project
	newProject.ID = 2
	newProject.APIKey = "APIKey2"
	newProject.UserID = suite.Project.UserID
	newProject.User = suite.User

	projectId, err := suite.Service.GetProjectService().AddProject(newProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	projects, err := suite.Service.GetProjectService().GetProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 2)
	assert.Equal(t, projects[1].ID, newProject.ID)
}

func TestUpdateProject(t *testing.T) {
	suite := tests.Startup()

	newProject := suite.Project
	newProject.APIKey = "APIKey2"

	err := suite.Service.GetProjectService().UpdateProject(newProject)
	assert.Nil(t, err)

	projects, err := suite.Service.GetProjectService().GetProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 1)
	assert.Equal(t, projects[0].ID, newProject.ID)
	assert.Equal(t, projects[0].APIKey, newProject.APIKey)
}

func TestDeleteProject(t *testing.T) {
	suite := tests.Startup()

	err := suite.Service.GetProjectService().DeleteProject(suite.Project.ID, models.User{})
	assert.NotNil(t, err)

	err = suite.Service.GetProjectService().DeleteProject(2, suite.User)
	assert.NotNil(t, err)

	err = suite.Service.GetProjectService().DeleteProject(suite.Project.ID, suite.User)
	assert.Nil(t, err)

	projects, err := suite.Service.GetProjectService().GetProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 0)

	err = suite.Service.GetProjectService().DeleteProject(suite.Project.ID, suite.User)
	assert.NotNil(t, err)
}
