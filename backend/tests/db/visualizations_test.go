package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"trackr/src/models"
	"trackr/tests"
)

func TestGetVisualizations(t *testing.T) {
	suite := tests.Startup()

	newProject := suite.Project
	newProject.ID = 2
	newProject.APIKey = "APIKey2"
	newProject.UserID = suite.Project.UserID
	newProject.User = suite.User

	projectId, err := suite.Service.GetProjectService().AddProject(newProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	visualizations, err := suite.Service.GetVisualizationService().GetVisualizations(models.Project{}, models.User{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(visualizations))

	visualizations, err = suite.Service.GetVisualizationService().GetVisualizations(suite.Project, models.User{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(visualizations))

	visualizations, err = suite.Service.GetVisualizationService().GetVisualizations(models.Project{}, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(visualizations))

	visualizations, err = suite.Service.GetVisualizationService().GetVisualizations(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(visualizations))
	assert.Equal(t, suite.Visualization.ID, visualizations[0].ID)
}

func TestGetVisualization(t *testing.T) {
	suite := tests.Startup()

	visualization, err := suite.Service.GetVisualizationService().GetVisualization(suite.Project.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)
	assert.Equal(t, suite.Visualization.ID, visualization.ID)

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(models.Project{}.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, visualization)

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(suite.Project.ID, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, visualization)

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(models.Project{}.ID, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, visualization)
}

func TestAddVisualization(t *testing.T) {
	suite := tests.Startup()

	newVisualization := suite.Visualization
	newVisualization.ID = 2
	newVisualization.Metadata = "Metadata"

	visualizationId, err := suite.Service.GetVisualizationService().AddVisualization(newVisualization)
	assert.Nil(t, err)
	assert.Equal(t, newVisualization.ID, visualizationId)

	visualizations, err := suite.Service.GetVisualizationService().GetVisualizations(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(visualizations), 2)
	assert.Equal(t, uint(2), visualizations[1].ID)
	assert.Equal(t, suite.Field.ID, visualizations[1].FieldID)
}

func TestUpdateVisualization(t *testing.T) {
	suite := tests.Startup()

	visualization := suite.Visualization
	visualization.Metadata = "Updated"

	err := suite.Service.GetVisualizationService().UpdateVisualization(visualization)
	assert.Nil(t, err)

	visualizations, err := suite.Service.GetVisualizationService().GetVisualizations(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(visualizations))
	assert.Equal(t, suite.Visualization.ID, visualizations[0].ID)
	assert.Equal(t, visualization.Metadata, visualizations[0].Metadata)
}

func TestDeleteVisualization(t *testing.T) {
	suite := tests.Startup()

	err := suite.Service.GetVisualizationService().DeleteVisualization(models.Visualization{})
	assert.NotNil(t, err)

	err = suite.Service.GetVisualizationService().DeleteVisualization(suite.Visualization)
	assert.Nil(t, err)

	visualizations, err := suite.Service.GetVisualizationService().GetVisualizations(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(visualizations))
}
