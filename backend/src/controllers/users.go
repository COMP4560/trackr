package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/services"
)

func getUserRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	numberOfFields, err := serviceProvider.GetFieldService().GetNumberOfFieldsByUser(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get number of fields."})
		return
	}

	numberOfValues, err := serviceProvider.GetValueService().GetNumberOfValuesByUser(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get number of values."})
		return
	}

	c.JSON(http.StatusOK, responses.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,

		NumberOfFields: numberOfFields,
		NumberOfValues: numberOfValues,

		MaxValueInterval: user.MaxValueInterval,
		MaxValues:        user.MaxValues,
	})
}

func updateUserRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	var json requests.UpdateUser
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	if json.FirstName == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Your first name cannot be empty."})
		return
	}

	if json.LastName == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Your last name cannot be empty."})
		return
	}

	user.FirstName = json.FirstName
	user.LastName = json.LastName

	if json.CurrentPassword != "" && json.NewPassword != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.CurrentPassword)); err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{Error: "Incorrect current password."})
			return
		}

		newHashedPassword, err := generateHashedPassword(json.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate new hashed password."})
			return
		}

		user.Password = newHashedPassword
	}

	user.UpdatedAt = time.Now()

	if err := serviceProvider.GetUserService().UpdateUser(*user); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to update user."})
		return
	}

	err := serviceProvider.GetLogService().AddLog("Updated user settings.", *user, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func deleteUserRoute(c *gin.Context) {
	user := getLoggedInUser(c)

	if err := serviceProvider.GetUserService().DeleteUser(*user); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete user."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func initUsersController(routerGroup *gin.RouterGroup, serviceProviderInput services.ServiceProvider, sessionMiddleware gin.HandlerFunc) {
	serviceProvider = serviceProviderInput

	usersRouterGroup := routerGroup.Group("/users")
	usersRouterGroup.Use(sessionMiddleware)

	usersRouterGroup.GET("/", getUserRoute)
	usersRouterGroup.PUT("/", updateUserRoute)
	usersRouterGroup.DELETE("/", deleteUserRoute)
}
