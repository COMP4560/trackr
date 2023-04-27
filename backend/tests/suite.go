package tests

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
	"trackr/src/controllers"
	"trackr/src/models"
	"trackr/src/services"
	"trackr/src/services_impl"
)

type Suite struct {
	Router  *gin.Engine
	Service services.ServiceProvider

	User           models.User
	Project        models.Project
	UserProject    models.UserProject
	Session        models.Session
	ExpiredSession models.Session
	Time           time.Time
	Field          models.Field
	Value          models.Value
	Visualization  models.Visualization
	Logs           []models.Log
}

func Startup() *Suite {
	var suite Suite

	suite.Service = services_impl.InitServiceProvider(sqlite.Open(":memory:?_foreign_keys=on"))
	suite.Time = time.Now()
	suite.User = models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email:            "Email@email",
		Password:         "$2a$12$Z4Ko/2d/EfenK9nBtpBRVO8I/3yOPnpcT/D/sbueRmhVDujVjHT4S",
		FirstName:        "FirstName",
		LastName:         "LastName",
		IsVerified:       true,
		MaxValues:        1,
		MaxValueInterval: 5,
	}
	suite.Service.GetUserService().AddUser(suite.User)
	savedUser, _ := suite.Service.GetUserService().GetUser(suite.User.Email)
	suite.User = *savedUser

	suite.Project = models.Project{
		Model: gorm.Model{
			ID: 1,
		},
		Name:        "Name",
		Description: "Description",
	}

	suite.UserProject = models.UserProject{
		ProjectID: 1,
		UserID:    suite.User.ID,
		Project:   suite.Project,
		User:      suite.User,
		Role:      "project_owner",
		DeletedAt: sql.NullTime{},
		APIKey:    "APIKey",
	}

	id, _ := suite.Service.GetProjectService().AddProject(suite.Project, suite.UserProject)
	userProject, _ := suite.Service.GetProjectService().GetUserProject(id, suite.User)
	suite.UserProject = *userProject
	suite.Project = userProject.Project

	suite.Session = models.Session{
		ID:        "SessionID",
		ExpiresAt: suite.Time.AddDate(1, 0, 0),

		UserID: suite.User.ID,
		User:   suite.User,
	}
	suite.Service.GetSessionService().AddSession(suite.Session)

	suite.ExpiredSession = models.Session{
		ID:        "ExpiredSessionID",
		ExpiresAt: suite.Time,

		UserID: suite.User.ID,
		User:   suite.User,
	}
	suite.Service.GetSessionService().AddSession(suite.ExpiredSession)

	suite.Field = models.Field{
		ID:        1,
		Name:      "Field1",
		ProjectID: suite.Project.ID,
		Project:   suite.Project,
		CreatedAt: time.Now(),
	}
	suite.Service.GetFieldService().AddField(suite.Field)

	suite.Visualization = models.Visualization{
		ID:       1,
		Metadata: "Metadata",

		FieldID:   suite.Field.ID,
		Field:     suite.Field,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	suite.Service.GetVisualizationService().AddVisualization(suite.Visualization)

	suite.Value = models.Value{
		ID:        1,
		Value:     "1.00",
		CreatedAt: suite.Time.Add(-time.Second * time.Duration(suite.User.MaxValueInterval)),

		FieldID: suite.Field.ID,
		Field:   suite.Field,
	}
	suite.Service.GetValueService().AddValue(suite.Value)

	suite.Service.GetLogService().AddLog("First Log", suite.User, nil)
	time.Sleep(time.Second)
	suite.Service.GetLogService().AddLog("Second Log", suite.User, &suite.Project.ID)
	suite.Logs = []models.Log{
		{
			Model: gorm.Model{
				ID: 2,
			},
			Message: "Second Log",

			ProjectID: &suite.Project.ID,
			Project:   suite.Project,

			UserID: suite.User.ID,
			User:   suite.User,
		},
		{
			Model: gorm.Model{
				ID: 1,
			},
			Message: "First Log",

			ProjectID: nil,
			Project:   models.Project{},

			UserID: suite.User.ID,
			User:   suite.User,
		},
	}

	return &suite
}

func StartupWithRouter(t *testing.T) *Suite {
	t.Setenv("DOCKER_ADDRESS", "172.18.0.2")
	t.Setenv("LOCAL_ADDRESS", "127.0.0.1")

	gin.SetMode(gin.ReleaseMode)

	suite := Startup()
	suite.Router = controllers.InitRouter(suite.Service)

	return suite
}
