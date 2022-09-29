package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trackr/src/models"
	"trackr/tests"
)

func TestGetFields(t *testing.T) {
	suite := tests.Startup()

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

func TestGetField(t *testing.T) {
	suite := tests.Startup()

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

	err := suite.Service.GetFieldService().AddField(newField)
	assert.Nil(t, err)

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

	err := suite.Service.GetFieldService().DeleteField(models.Field{}.ID, suite.User)
	assert.NotNil(t, err)

	err = suite.Service.GetFieldService().DeleteField(suite.Field.ID, models.User{})
	assert.NotNil(t, err)

	err = suite.Service.GetFieldService().DeleteField(suite.Field.ID, suite.User)
	assert.Nil(t, err)

	fields, err := suite.Service.GetFieldService().GetFields(suite.Project, suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(fields))
}
