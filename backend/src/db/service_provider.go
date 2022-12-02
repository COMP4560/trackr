package db

import (
	"gorm.io/gorm"

	"trackr/src/models"
	"trackr/src/services"
)

type ServiceProviderDB struct {
	sessionService       services.SessionService
	userService          services.UserService
	projectService       services.ProjectService
	fieldService         services.FieldService
	visualizationService services.VisualizationService
	valueService         services.ValueService
	logService           services.LogService
}

func InitServiceProvider(dialector gorm.Dialector) services.ServiceProvider {
	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil
	}

	database.Exec("PRAGMA foreign_keys = ON", nil)
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Session{})
	database.AutoMigrate(&models.Project{})
	database.AutoMigrate(&models.Field{})
	database.AutoMigrate(&models.Value{})
	database.AutoMigrate(&models.Visualization{})
	database.AutoMigrate(&models.Log{})

	serviceProviderDB := &ServiceProviderDB{}
	serviceProviderDB.sessionService = &SessionServiceDB{database: database}
	serviceProviderDB.userService = &UserServiceDB{database: database}
	serviceProviderDB.projectService = &ProjectServiceDB{database: database}
	serviceProviderDB.fieldService = &FieldServiceDB{database: database}
	serviceProviderDB.valueService = &ValueServiceDB{database: database}
	serviceProviderDB.visualizationService = &VisulizationServiceDB{database: database}
	serviceProviderDB.logService = &LogServiceDB{database: database}

	return serviceProviderDB
}

func (serviceProviderDB *ServiceProviderDB) GetSessionService() services.SessionService {
	return serviceProviderDB.sessionService
}

func (serviceProviderDB *ServiceProviderDB) GetUserService() services.UserService {
	return serviceProviderDB.userService
}

func (serviceProviderDB *ServiceProviderDB) GetProjectService() services.ProjectService {
	return serviceProviderDB.projectService
}

func (serviceProviderDB *ServiceProviderDB) GetFieldService() services.FieldService {
	return serviceProviderDB.fieldService
}

func (serviceProviderDB *ServiceProviderDB) GetVisualizationService() services.VisualizationService {
	return serviceProviderDB.visualizationService
}

func (serviceProviderDB *ServiceProviderDB) GetValueService() services.ValueService {
	return serviceProviderDB.valueService
}

func (serviceProviderDB *ServiceProviderDB) GetLogService() services.LogService {
	return serviceProviderDB.logService
}
