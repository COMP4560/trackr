package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"trackr/src/models"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/services"
)

func addFieldRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	var json requests.AddField
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	userProject, err := serviceProvider.GetProjectService().GetUserProject(json.ProjectID, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Cannot find project."})
		return
	}

	if json.Name == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "The name of a field cannot be empty."})
		return
	}

	field := models.Field{
		Name:      json.Name,
		Project:   userProject.Project,
		CreatedAt: time.Now(),
	}

	field.ID, err = serviceProvider.GetFieldService().AddField(field)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a new field."})
		return
	}

	err = serviceProvider.GetLogService().AddLog(fmt.Sprintf("Added the field %s.", field.Name), *user, &userProject.ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.NewField{
		ID: field.ID,
	})
}

func getFieldsRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid :projectId parameter provided."})
		return
	}

	userProject, err := serviceProvider.GetProjectService().GetUserProject(uint(projectId), *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find project."})
		return
	}

	fields, err := serviceProvider.GetFieldService().GetFields(userProject.Project, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to get fields."})
		return
	}

	fieldList := make([]responses.Field, len(fields))
	for index, field := range fields {
		numberOfValues, err := serviceProvider.GetValueService().GetNumberOfValuesByField(field)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to get number of values."})
			return
		}

		fieldList[index] = responses.Field{
			ID:             field.ID,
			Name:           field.Name,
			CreatedAt:      field.CreatedAt,
			NumberOfValues: numberOfValues,
		}
	}

	c.JSON(http.StatusOK, responses.FieldList{Fields: fieldList})
}

func updateFieldRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	var json requests.UpdateField
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	field, err := serviceProvider.GetFieldService().GetField(json.ID, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find field."})
		return
	}

	if json.Name == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "The name of a field cannot be empty."})
		return
	}

	field.Name = json.Name
	field.UpdatedAt = time.Now()

	if err := serviceProvider.GetFieldService().UpdateField(*field); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to update field."})
		return
	}

	err = serviceProvider.GetLogService().AddLog(fmt.Sprintf("Modified the field %s.", field.Name), *user, &field.ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func deleteFieldRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	fieldId, err := strconv.Atoi(c.Param("fieldId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid :fieldId parameter provided."})
		return
	}

	field, err := serviceProvider.GetFieldService().GetField(uint(fieldId), *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to get field."})
		return
	}

	err = serviceProvider.GetFieldService().DeleteField(*field)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete field."})
		return
	}

	err = serviceProvider.GetLogService().AddLog(fmt.Sprintf("Deleted the field %s.", field.Name), *user, &field.ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func initFieldsController(routerGroup *gin.RouterGroup, serviceProviderInput services.ServiceProvider, sessionMiddleware gin.HandlerFunc) {
	serviceProvider = serviceProviderInput

	fieldsRouterGroup := routerGroup.Group("/fields")
	fieldsRouterGroup.Use(sessionMiddleware)

	fieldsRouterGroup.POST("/", addFieldRoute)
	fieldsRouterGroup.GET("/:projectId", getFieldsRoute)
	fieldsRouterGroup.PUT("/", updateFieldRoute)
	fieldsRouterGroup.DELETE("/:fieldId", deleteFieldRoute)
}
