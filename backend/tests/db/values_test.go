package db_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"

	"trackr/src/models"
	"trackr/tests"
)

func TestGetValues(t *testing.T) {
	suite := tests.Startup()

	newProject := suite.Project
	newProject.ID = 2
	newProject.APIKey = "APIKey2"
	newProject.UserID = suite.Project.UserID
	newProject.User = suite.User

	projectId, err := suite.Service.GetProjectService().AddProject(newProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	values, err := suite.Service.GetValueService().GetValues(models.Field{}, models.User{}, "asc", 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(values))

	values, err = suite.Service.GetValueService().GetValues(suite.Field, models.User{}, "asc", 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(values))

	values, err = suite.Service.GetValueService().GetValues(models.Field{}, suite.User, "asc", 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(values))

	expectedValues := []models.Value{suite.Value}
	numberOfExpectedValues := 10 + len(expectedValues)

	for i := 2; i <= numberOfExpectedValues; i++ {
		newValue := suite.Value
		newValue.ID = uint(i)
		newValue.CreatedAt = time.Now()
		newValue.Value = fmt.Sprintf("%d.00", i)

		err := suite.Service.GetValueService().AddValue(newValue)
		assert.Nil(t, err)

		numberOfValues, err := suite.Service.GetValueService().GetNumberOfValuesByUser(suite.User)
		assert.Nil(t, err)
		assert.Equal(t, int64(i), numberOfValues)

		numberOfValues, err = suite.Service.GetValueService().GetNumberOfValuesByField(suite.Field)
		assert.Nil(t, err)
		assert.Equal(t, int64(i), numberOfValues)

		expectedValues = append(expectedValues, newValue)
	}

	for offset := 0; offset < numberOfExpectedValues; offset++ {
		for limit := 0; limit <= numberOfExpectedValues; limit++ {
			for _, order := range []string{"desc", "asc"} {
				expectedValuesSlice := expectedValues[:]

				sort.Slice(expectedValuesSlice, func(i, j int) bool {
					if order == "desc" {
						return expectedValuesSlice[i].CreatedAt.After(
							expectedValuesSlice[j].CreatedAt,
						)
					} else {
						return expectedValuesSlice[i].CreatedAt.Before(
							expectedValuesSlice[j].CreatedAt,
						)
					}
				})

				if offset == 0 {
					expectedValuesSlice = expectedValuesSlice[:]
				} else {
					expectedValuesSlice = expectedValuesSlice[offset:]
				}

				values, err = suite.Service.GetValueService().GetValues(
					suite.Field, suite.User, order, offset, limit,
				)
				assert.Nil(t, err)

				if limit > 0 {
					assert.LessOrEqual(t, len(values), limit)
				} else {
					assert.Equal(t, len(expectedValuesSlice), len(values))
				}

				for i := 0; i < len(values); i++ {
					assert.Equal(t, expectedValuesSlice[i].Value, values[i].Value)
				}
			}
		}
	}
}

func TestGetValue(t *testing.T) {
	suite := tests.Startup()

	value, err := suite.Service.GetValueService().GetValue(models.Value{}.ID, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = suite.Service.GetValueService().GetValue(models.Value{}.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = suite.Service.GetValueService().GetValue(suite.Value.ID, models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, suite.Value.Value, value.Value)
}

func TestGetLastAddedValue(t *testing.T) {
	suite := tests.Startup()

	value, err := suite.Service.GetValueService().GetLastAddedValue(models.User{})
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = suite.Service.GetValueService().GetLastAddedValue(suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, suite.Value.Value, value.Value)

	newValue := suite.Value
	newValue.ID = 2
	newValue.Value = "2.00"
	newValue.CreatedAt = time.Now().Add(time.Minute)

	err = suite.Service.GetValueService().AddValue(newValue)
	assert.Nil(t, err)

	value, err = suite.Service.GetValueService().GetLastAddedValue(suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, newValue.Value, value.Value)

	err = suite.Service.GetValueService().DeleteValues(newValue.Field)
	assert.Nil(t, err)

	value, err = suite.Service.GetValueService().GetLastAddedValue(suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestAddValue(t *testing.T) {
	suite := tests.Startup()

	newValue := suite.Value
	newValue.ID = 2
	newValue.Value = "2.00"

	value, err := suite.Service.GetValueService().GetValue(newValue.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	err = suite.Service.GetValueService().AddValue(newValue)
	assert.Nil(t, err)

	value, err = suite.Service.GetValueService().GetValue(newValue.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, newValue.ID, value.ID)
	assert.Equal(t, newValue.Value, value.Value)
}

func TestDeleteValues(t *testing.T) {
	suite := tests.Startup()

	newValue := suite.Value
	newValue.ID = 2
	newValue.Value = "2.00"

	err := suite.Service.GetValueService().AddValue(newValue)
	assert.Nil(t, err)

	value, err := suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, suite.Value.ID, value.ID)

	value, err = suite.Service.GetValueService().GetValue(newValue.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, newValue)
	assert.Equal(t, newValue.ID, value.ID)

	err = suite.Service.GetValueService().DeleteValues(suite.Field)
	assert.Nil(t, err)

	value, err = suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = suite.Service.GetValueService().GetValue(newValue.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestGetNumberOfValuesByUser(t *testing.T) {
	suite := tests.Startup()

	numberOfValues, err := suite.Service.GetValueService().GetNumberOfValuesByUser(models.User{})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfValues)

	newValue := suite.Value
	newValue.ID = 2
	newValue.Value = "2.00"

	err = suite.Service.GetValueService().AddValue(newValue)
	assert.Nil(t, err)

	numberOfValues, err = suite.Service.GetValueService().GetNumberOfValuesByUser(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), numberOfValues)

	err = suite.Service.GetValueService().DeleteValues(suite.Field)
	assert.Nil(t, err)

	numberOfValues, err = suite.Service.GetValueService().GetNumberOfValuesByUser(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfValues)
}

func TestGetNumberOfValuesByField(t *testing.T) {
	suite := tests.Startup()

	numberOfValues, err := suite.Service.GetValueService().GetNumberOfValuesByField(models.Field{})
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfValues)

	newValue := suite.Value
	newValue.ID = 2
	newValue.Value = "2.00"

	err = suite.Service.GetValueService().AddValue(newValue)
	assert.Nil(t, err)

	numberOfValues, err = suite.Service.GetValueService().GetNumberOfValuesByField(suite.Field)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), numberOfValues)

	err = suite.Service.GetValueService().DeleteValues(suite.Field)
	assert.Nil(t, err)

	numberOfValues, err = suite.Service.GetValueService().GetNumberOfValuesByField(suite.Field)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfValues)
}
