package helpers

import (
	"trackr/src/models"
	"trackr/tests"
)

func CreateNewProject(suite tests.Suite) (models.Project, models.UserProject) {
	newProject := suite.Project
	newProject.ID = 2

	userProject := suite.UserProject
	userProject.Project = newProject
	userProject.User = suite.User
	userProject.APIKey = "APIKey2"

	return newProject, userProject
}
