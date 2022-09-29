package controllers

import (
	"github.com/gin-gonic/gin"

	"trackr/src/services"
)

var (
	serviceProvider services.ServiceProvider
)

func InitRouter(serviceProviderInput services.ServiceProvider) *gin.Engine {
	serviceProvider = serviceProviderInput

	router := gin.Default()
	routerGroup := router.Group("/api")
	sessionMiddleware := initAuthMiddleware(serviceProvider)

	initAuthController(routerGroup, serviceProvider)
	initProjectsController(routerGroup, serviceProvider, sessionMiddleware)
	initUsersController(routerGroup, serviceProvider, sessionMiddleware)
	initFieldsController(routerGroup, serviceProvider, sessionMiddleware)

	return router
}
