package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/models"
	"trackr/src/services"
)

func addFieldRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	var json requests.AddField
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	project, err := serviceProvider.GetProjectService().GetProjectByIdAndUser(json.ProjectID, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Cannot find project."})
		return
	}

	if json.Name == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "The name of a field cannot be empty."})
		return
	}

	createdAt := time.Now()

	field := models.Field{
		Name:      json.Name,
		UpdatedAt: createdAt,
		CreatedAt: createdAt,

		Project: *project,
	}

	if err := serviceProvider.GetFieldService().AddField(field); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a new field."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func getFieldsRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid :projectId parameter provided."})
		return
	}

	project, err := serviceProvider.GetProjectService().GetProjectByIdAndUser(uint(projectId), *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find project."})
		return
	}

	fields, err := serviceProvider.GetFieldService().GetFields(*project, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to get fields."})
		return
	}

	fieldList := make([]responses.Field, len(fields))
	for index, field := range fields {
		fieldList[index] = responses.Field{
			ID:   field.ID,
			Name: field.Name,
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

	wasModified := false

	if json.Name != "" {
		field.Name = json.Name
		wasModified = true
	}

	if wasModified {
		field.UpdatedAt = time.Now()
	}

	if err := serviceProvider.GetFieldService().UpdateField(*field); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to update field."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func deleteFieldRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	fieldId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid :id parameter provided."})
		return
	}

	err = serviceProvider.GetFieldService().DeleteField(uint(fieldId), *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete field."})
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
	fieldsRouterGroup.DELETE("/:id", deleteFieldRoute)
}
