package tests

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"

	"trackr/src/controllers"
	"trackr/src/db"
	"trackr/src/models"
	"trackr/src/services"
)

type Suite struct {
	Router  *gin.Engine
	Service services.ServiceProvider

	User           models.User
	Project        models.Project
	Session        models.Session
	ExpiredSession models.Session
	Time           time.Time
	Field          models.Field
}

func Startup() *Suite {
	var suite Suite

	suite.Service = db.InitServiceProvider(sqlite.Open(":memory:"))
	suite.Time = time.Now()
	suite.User = models.User{
		ID:          1,
		Email:       "Email@email",
		Password:    "$2a$12$Z4Ko/2d/EfenK9nBtpBRVO8I/3yOPnpcT/D/sbueRmhVDujVjHT4S",
		FirstName:   "FirstName",
		LastName:    "LastName",
		UpdatedAt:   suite.Time,
		CreatedAt:   suite.Time,
		IsVerified:  true,
		MaxValues:   1,
		MaxProjects: 2,
	}
	suite.Service.GetUserService().AddUser(suite.User)

	suite.Project = models.Project{
		ID:          1,
		Name:        "Name",
		Description: "Description",
		APIKey:      "APIKey",
		ShareURL:    nil,
		CreatedAt:   suite.Time,
		UpdatedAt:   suite.Time,

		UserID: suite.User.ID,
		User:   suite.User,
	}
	suite.Service.GetProjectService().AddProject(suite.Project)

	suite.Session = models.Session{
		ID:        "SessionID",
		CreatedAt: suite.Time,
		ExpiresAt: suite.Time.AddDate(1, 0, 0),

		UserID: suite.User.ID,
		User:   suite.User,
	}
	suite.Service.GetSessionService().AddSession(suite.Session)

	suite.ExpiredSession = models.Session{
		ID:        "ExpiredSessionID",
		CreatedAt: suite.Time,
		ExpiresAt: suite.Time,

		UserID: suite.User.ID,
		User:   suite.User,
	}
	suite.Service.GetSessionService().AddSession(suite.ExpiredSession)

	suite.Field = models.Field{
		ID:        1,
		Name:      "Field1",
		UpdatedAt: suite.Time,
		CreatedAt: suite.Time,

		ProjectID: suite.Project.ID,
		Project:   suite.Project,
	}
	suite.Service.GetFieldService().AddField(suite.Field)

	return &suite
}

func StartupWithRouter() *Suite {
	gin.SetMode(gin.ReleaseMode)

	suite := Startup()
	suite.Router = controllers.InitRouter(suite.Service)

	return suite
}
