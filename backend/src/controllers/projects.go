package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trackr/src/forms/responses/projects"
	"trackr/src/models"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/services"
)

const (
	apiKeyLength   = 64
	shareURLLength = 32
)

func addProjectRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	apiKey, err := generateAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate API key."})
		return
	}

	project := models.Project{
		Name:        "Untitled Project",
		Description: "",
		// OrganizationID: todo,
	}

	userProject := models.UserProject{
		UserID: user.ID,
		Role:   "organization_owner",
		APIKey: apiKey,
	}

	project.ID, err = serviceProvider.GetProjectService().AddProject(project, userProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a new project."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Created a new project.", *user, &project.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, projects.NewProject{
		ID: project.ID,
	})
}

func getProjectRoute(c *gin.Context) {
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

	project := userProject.Project
	numberOfFields, err := serviceProvider.GetFieldService().GetNumberOfFieldsByProject(project, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get number of fields."})
		return
	}

	c.JSON(http.StatusOK, projects.Project{
		ID:             project.ID,
		Name:           project.Name,
		Description:    project.Description,
		APIKey:         userProject.APIKey,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
		NumberOfFields: numberOfFields,
	})
}

func getProjectsRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	userProjects, err := serviceProvider.GetProjectService().GetUserProjects(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get projects."})
		return
	}

	projectList := make([]projects.Project, len(userProjects))
	for index, userProject := range userProjects {
		project := userProject.Project

		numberOfFields, err := serviceProvider.GetFieldService().GetNumberOfFieldsByProject(project, *user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get number of fields."})
			return
		}

		projectList[index] = projects.Project{
			ID:             project.ID,
			Name:           project.Name,
			Description:    project.Description,
			APIKey:         userProject.APIKey,
			CreatedAt:      project.CreatedAt,
			UpdatedAt:      project.UpdatedAt,
			NumberOfFields: numberOfFields,
		}
	}

	c.JSON(http.StatusOK, projects.ProjectList{Projects: projectList})
}

func deleteProjectRoute(c *gin.Context) {
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

	err = serviceProvider.GetProjectService().DeleteProject(uint(projectId), *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete project."})
		return
	}

	err = serviceProvider.GetLogService().AddLog(fmt.Sprintf("Deleted the project %s.", userProject.Project.Name), *user, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func updateProjectRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	var json requests.UpdateProject
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	userProject, err := serviceProvider.GetProjectService().GetUserProject(json.ID, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Failed to find project."})
		return
	}

	if json.Name == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "The project's name cannot be empty."})
		return
	}

	project := userProject.Project
	project.Name = json.Name
	project.Description = json.Description

	if json.ResetAPIKey {
		apiKey, err := generateAPIKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate new API key."})
			return
		}

		userProject.APIKey = apiKey
	}

	err = serviceProvider.GetProjectService().UpdateProject(project, *userProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to update project."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Updated the project's information.", *user, &project.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, projects.UpdateProject{
		APIKey: userProject.APIKey,
	})
}

func initProjectsController(routerGroup *gin.RouterGroup, serviceProviderInput services.ServiceProvider, sessionMiddleware gin.HandlerFunc) {
	serviceProvider = serviceProviderInput

	projectsRouterGroup := routerGroup.Group("/projects")
	projectsRouterGroup.Use(sessionMiddleware)

	projectsRouterGroup.POST("/", addProjectRoute)
	projectsRouterGroup.GET("/:projectId", getProjectRoute)
	projectsRouterGroup.GET("/", getProjectsRoute)
	projectsRouterGroup.PUT("/", updateProjectRoute)
	projectsRouterGroup.DELETE("/:projectId", deleteProjectRoute)
}
