package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/forms/responses/organizations"
	"trackr/src/models"
	"trackr/src/services"
)

func addOrganizationRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	apiKey, err := generateAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate API key."})
		return
	}

	organization := models.Organization{
		Name:        "Untitled Organization",
		Description: "",
	}

	userOrganization := models.UserOrganization{
		UserID: user.ID,
		Role:   "organization_owner",
		APIKey: apiKey,
	}

	organization.ID, err = serviceProvider.GetOrganizationService().AddOrganization(organization, userOrganization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a new organization."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Created a new organization.", *user, &organization.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, organizations.NewOrganization{
		ID: organization.ID,
	})
}

func getOrganizationRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	organizationId, err := strconv.Atoi(c.Param("organizationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid :organizationId parameter provided."})
		return
	}

	userOrganization, err := serviceProvider.GetOrganizationService().GetUserOrganization(uint(organizationId), *user)
	organization := userOrganization.Organization
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find organization."})
		return
	}

	numberOfProjects, err := serviceProvider.GetFieldService().GetNumberOfProjectsByOrganization(organization, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get number of projects."})
		return
	}

	c.JSON(http.StatusOK, organizations.Organization{
		ID:               organization.ID,
		Name:             organization.Name,
		Description:      organization.Description,
		APIKey:           userOrganization.APIKey,
		CreatedAt:        organization.CreatedAt,
		UpdatedAt:        organization.UpdatedAt,
		NumberOfProjects: numberOfProjects,
	})
}

func getOrganizationsRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	userOrganizations, err := serviceProvider.GetOrganizationService().GetUserOrganizations(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get organizations."})
		return
	}

	organizationList := make([]organizations.Organization, len(userOrganizations))
	for index, organization := range userOrganizations {
		organizationEntity := organization.Organization
		numberOfProjects, err := serviceProvider.GetFieldService().GetNumberOfProjectsByOrganization(organizationEntity, *user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get number of fields."})
			return
		}

		organizationList[index] = organizations.Organization{
			ID:               organizationEntity.ID,
			Name:             organizationEntity.Name,
			Description:      organizationEntity.Description,
			APIKey:           organization.APIKey,
			CreatedAt:        organizationEntity.CreatedAt,
			UpdatedAt:        organizationEntity.UpdatedAt,
			NumberOfProjects: numberOfProjects,
		}
	}

	c.JSON(http.StatusOK, organizations.OrganizationList{Organizations: organizationList})
}

func deleteOrganizationRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	organizationId, err := strconv.Atoi(c.Param("organizationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid :organizationId parameter provided."})
		return
	}

	userOrganization, err := serviceProvider.GetOrganizationService().GetUserOrganization(uint(organizationId), *user)
	organization := userOrganization.Organization
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find organization."})
		return
	}

	err = serviceProvider.GetOrganizationService().DeleteOrganization(uint(organizationId), *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete organization."})
		return
	}

	err = serviceProvider.GetLogService().AddLog(fmt.Sprintf("Deleted the organization %s.", organization.Name), *user, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func updateOrganizationRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	var json requests.UpdateOrganization
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	userOrganization, err := serviceProvider.GetOrganizationService().GetUserOrganization(json.ID, *user)
	organization := userOrganization.Organization
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find organization."})
		return
	}

	if json.Name == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "The organization's name cannot be empty."})
		return
	}

	organization.Name = json.Name
	organization.Description = json.Description

	if json.ResetAPIKey {
		apiKey, err := generateAPIKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate new API key."})
			return
		}

		userOrganization.APIKey = apiKey
	}

	organization.UpdatedAt = time.Now()

	err = serviceProvider.GetOrganizationService().UpdateOrganization(organization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to update organization."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Updated the organization's information.", *user, &organization.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, organizations.UpdateOrganization{
		APIKey: userOrganization.APIKey,
	})
}

func initOrganizationsController(routerGroup *gin.RouterGroup, serviceProviderInput services.ServiceProvider, sessionMiddleware gin.HandlerFunc) {
	serviceProvider = serviceProviderInput

	organizationsRouterGroup := routerGroup.Group("/organizations")
	organizationsRouterGroup.Use(sessionMiddleware)

	organizationsRouterGroup.POST("/", addOrganizationRoute)
	organizationsRouterGroup.GET("/:organizationId", getOrganizationRoute)
	organizationsRouterGroup.GET("/", getOrganizationsRoute)
	organizationsRouterGroup.PUT("/", updateOrganizationRoute)
	organizationsRouterGroup.DELETE("/:organizationId", deleteOrganizationRoute)
}
