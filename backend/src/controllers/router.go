package controllers

import (
	"fmt"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"trackr/src/services"
)

var (
	serviceProvider services.ServiceProvider
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("API_PATH"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitRouter(serviceProviderInput services.ServiceProvider) *gin.Engine {
	serviceProvider = serviceProviderInput

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(corsMiddleware())

	router.ForwardedByClientIP = false
	err := router.SetTrustedProxies([]string{os.Getenv("DOCKER_ADDRESS"), os.Getenv("LOCAL_ADDRESS")})
	if err != nil {
		fmt.Println("Error setting trusted proxies:", err)
		return nil
	}

	routerGroup := router.Group("/api")
	sessionMiddleware := initAuthMiddleware(serviceProvider)

	initAuthController(routerGroup, serviceProvider)
	initOrganizationsController(routerGroup, serviceProvider, sessionMiddleware)
	initProjectsController(routerGroup, serviceProvider, sessionMiddleware)
	initUsersController(routerGroup, serviceProvider, sessionMiddleware)
	initFieldsController(routerGroup, serviceProvider, sessionMiddleware)
	initValuesController(routerGroup, serviceProvider, sessionMiddleware)
	initVisualizationsController(routerGroup, serviceProvider, sessionMiddleware)
	initLogsController(routerGroup, serviceProvider, sessionMiddleware)

	return router
}
