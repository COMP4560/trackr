package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"trackr/src/forms/responses"
	"trackr/src/services"
)

func getLogsRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	logs, err := serviceProvider.GetLogService().GetLogs(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get logs."})
		return
	}

	logList := make([]responses.Log, len(logs))
	for index, log := range logs {
		projectName := ""
		if log.ProjectID != nil {
			projectName = log.Project.Name
		}

		logList[index] = responses.Log{
			Message:   log.Message,
			CreatedAt: log.CreatedAt,

			ProjectID:   log.ProjectID,
			ProjectName: projectName,
		}
	}

	c.JSON(http.StatusOK, responses.LogList{Logs: logList})
}

func initLogsController(routerGroup *gin.RouterGroup, serviceProviderInput services.ServiceProvider, sessionMiddleware gin.HandlerFunc) {
	serviceProvider = serviceProviderInput

	usersRouterGroup := routerGroup.Group("/logs")
	usersRouterGroup.Use(sessionMiddleware)
	usersRouterGroup.GET("/", getLogsRoute)
}
