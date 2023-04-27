package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trackr/src/models"
	"trackr/tests"
	"trackr/tests/helpers"
)

func TestGetFields(t *testing.T) {
	suite := tests.Startup()

	newProject, newUserProject := helpers.CreateNewProject(*suite)

	projectId, err := suite.Service.GetProjectService().AddProject(newProject, newUserProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	fields, err := suite.Service.GetFieldService().GetFields(models.Project{}, models.User{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(fields))

	fields, err = suite.Service.GetFieldService().GetFields(suite.Project, models.User{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(fields))

	fields, err = suite.Service.GetFieldService().GetFields(models.Project{}, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(fields))

	fields, err = suite.Service.GetFieldService().GetFields(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fields))
	assert.Equal(t, suite.Field.ID, fields[0].ID)
}

func TestGetNumberOfFieldsByProject(t *testing.T) {
	suite := tests.Startup()

	numberOfFields, err := suite.Service.GetFieldService().GetNumberOfFieldsByProject(models.Project{}, models.User{})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfFields)

	numberOfFields, err = suite.Service.GetFieldService().GetNumberOfFieldsByProject(suite.Project, models.User{})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfFields)

	numberOfFields, err = suite.Service.GetFieldService().GetNumberOfFieldsByProject(models.Project{}, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfFields)

	numberOfFields, err = suite.Service.GetFieldService().GetNumberOfFieldsByProject(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), numberOfFields)
}

func TestGetNumberOfFieldsByUser(t *testing.T) {
	suite := tests.Startup()

	numberOfFields, err := suite.Service.GetFieldService().GetNumberOfFieldsByUser(models.User{})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfFields)

	numberOfFields, err = suite.Service.GetFieldService().GetNumberOfFieldsByUser(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), numberOfFields)

	newProject, newUserProject := helpers.CreateNewProject(*suite)

	projectId, err := suite.Service.GetProjectService().AddProject(newProject, newUserProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	newField := suite.Field
	newField.ID = 2
	newField.Name = "Field2"
	newField.CreatedAt = suite.Time
	newField.UpdatedAt = suite.Time
	newField.Project = newProject

	fieldId, err := suite.Service.GetFieldService().AddField(newField)
	assert.Nil(t, err)
	assert.Equal(t, newField.ID, fieldId)

	numberOfFields, err = suite.Service.GetFieldService().GetNumberOfFieldsByUser(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), numberOfFields)
}

func TestGetField(t *testing.T) {
	suite := tests.Startup()

	newProject, newUserProject := helpers.CreateNewProject(*suite)

	projectId, err := suite.Service.GetProjectService().AddProject(newProject, newUserProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	field, err := suite.Service.GetFieldService().GetField(models.Field{}.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, field)

	field, err = suite.Service.GetFieldService().GetField(suite.Field.ID, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, field)

	field, err = suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, field)
	assert.Equal(t, suite.Field.ID, field.ID)
}

func TestAddField(t *testing.T) {
	suite := tests.Startup()

	newField := suite.Field
	newField.ID = 2
	newField.Name = "Field2"
	newField.CreatedAt = suite.Time
	newField.UpdatedAt = suite.Time
	newField.Project = suite.Project

	fieldId, err := suite.Service.GetFieldService().AddField(newField)
	assert.Nil(t, err)
	assert.Equal(t, newField.ID, fieldId)

	fields, err := suite.Service.GetFieldService().GetFields(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(fields), 2)
	assert.Equal(t, uint(2), fields[1].ID)
	assert.Equal(t, suite.Project.ID, fields[1].ProjectID)
}

func TestUpdateField(t *testing.T) {
	suite := tests.Startup()

	newField := suite.Field
	newField.Name = "NewField"

	err := suite.Service.GetFieldService().UpdateField(newField)
	assert.Nil(t, err)

	fields, err := suite.Service.GetFieldService().GetFields(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fields))
	assert.Equal(t, suite.Field.ID, fields[0].ID)
	assert.Equal(t, newField.Name, fields[0].Name)
}

func TestDeleteField(t *testing.T) {
	suite := tests.Startup()

	value, err := suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)

	visualization, err := suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)

	err = suite.Service.GetFieldService().DeleteField(models.Field{})
	assert.NotNil(t, err)

	err = suite.Service.GetFieldService().DeleteField(suite.Field)
	assert.Nil(t, err)

	fields, err := suite.Service.GetFieldService().GetFields(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(fields))

	err = suite.Service.GetFieldService().DeleteField(suite.Field)
	assert.NotNil(t, err)

	value, err = suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, visualization)
}
