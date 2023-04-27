package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"
	"trackr/src/models"

	"trackr/src/common"
	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/services"
)

const (
	sessionIdLength = 64
	sessionCookie   = "Session"
)

func loginRoute(c *gin.Context) {
	if isLoggedIn(c) != nil {
		c.JSON(http.StatusTemporaryRedirect, responses.Error{Error: "You are currently logged in."})
		return
	}

	var json requests.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	if json.Email == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide an email."})
		return
	}

	if json.Password == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide a password."})
		return
	}

	user, err := serviceProvider.GetUserService().GetUser(json.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.Error{Error: "Invalid email or password combination."})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, responses.Error{Error: "Invalid email or password combination."})
		return
	}

	err = serviceProvider.GetSessionService().DeleteExpiredSessions(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.Error{Error: "Failed to clean up sessions."})
		return
	}

	sessionId, err := common.RandomString(sessionIdLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate session identifier."})
		return
	}

	expiresAt := time.Now()

	if json.RememberMe {
		expiresAt = expiresAt.AddDate(0, 1, 0)
	} else {
		expiresAt = expiresAt.AddDate(0, 0, 7)
	}

	session := models.Session{
		ID:        sessionId,
		UserAgent: c.Request.UserAgent(),
		ExpiresAt: expiresAt,
		User:      *user,
	}

	err = serviceProvider.GetSessionService().AddSession(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create session."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Signed in.", *user, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.SetCookie(sessionCookie, sessionId, math.MaxInt32, "/", "*", false, true)
	c.JSON(http.StatusOK, responses.Empty{})
}

func logoutRoute(c *gin.Context) {
	sessionId, err := c.Cookie(sessionCookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "No session cookie provided."})
		return
	}

	user := isLoggedIn(c)
	if user == nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You are currently not logged in."})
		return
	}

	err = serviceProvider.GetSessionService().DeleteExpiredSessions(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete expired sessions."})
		return
	}

	err = serviceProvider.GetSessionService().DeleteSession(sessionId, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to delete session."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Signed out.", *user, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func registerRoute(c *gin.Context) {
	if os.Getenv("DISABLE_SIGN_UP") == "true" {
		c.JSON(http.StatusForbidden, responses.Error{Error: "Registration is currently disabled. Ask your administrator to enable it."})
		return
	}

	if isLoggedIn(c) != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You cannot make a new account while you are currently logged in."})
		return
	}

	var json requests.Register
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Invalid request parameters provided."})
		return
	}

	if json.FirstName == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide a first name."})
		return
	}

	if json.LastName == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide a last name."})
		return
	}

	if json.Email == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide an email address."})
		return
	}

	if json.Password == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide a password."})
		return
	}

	if _, err := mail.ParseAddress(json.Email); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "You must provide a valid email address."})
		return
	}

	count, err := serviceProvider.GetUserService().GetNumberOfUsers(json.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to get user count."})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "The email is already taken."})
		return
	}

	hashedPassword, err := generateHashedPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to generate password."})
		return
	}

	maxValues, _ := strconv.ParseInt(os.Getenv("MAX_VALUES"), 10, 64)
	maxValueInterval, _ := strconv.ParseInt(os.Getenv("MAX_VALUE_INTERVAL"), 10, 64)

	user := models.User{
		Email:      json.Email,
		Password:   hashedPassword,
		FirstName:  json.FirstName,
		LastName:   json.LastName,
		IsVerified: false,

		MaxValues:        maxValues,
		MaxValueInterval: maxValueInterval,
	}

	user.ID, err = serviceProvider.GetUserService().AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to register new user."})
		return
	}

	err = serviceProvider.GetLogService().AddLog("Created a new account.", user, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "Failed to create a log entry."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func isLoggedInRoute(c *gin.Context) {
	if isLoggedIn(c) == nil {
		c.JSON(http.StatusUnauthorized, responses.Error{Error: "Not logged in."})
		return
	}

	c.JSON(http.StatusOK, responses.Empty{})
}

func initAuthMiddleware(serviceProviderInput services.ServiceProvider) gin.HandlerFunc {
	serviceProvider = serviceProviderInput

	return func(c *gin.Context) {
		user := isLoggedIn(c)
		if user == nil {
			c.JSON(http.StatusForbidden, responses.Error{Error: "Not authorized to access this resource."})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
	}
}

func initAuthController(routerGroup *gin.RouterGroup, serviceProviderInput services.ServiceProvider) {
	serviceProvider = serviceProviderInput

	authRouterGroup := routerGroup.Group("/auth")
	authRouterGroup.POST("/login", loginRoute)
	authRouterGroup.GET("/", isLoggedInRoute)
	authRouterGroup.GET("/logout", logoutRoute)
	authRouterGroup.POST("/register", registerRoute)
}
