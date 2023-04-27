package services_impl

import (
	"fmt"
	"gorm.io/gorm"
	"trackr/src/models"
	"trackr/src/services"
	"trackr/src/services_impl/db"
)

type ServiceProvider struct {
	sessionService       services.SessionService
	userService          services.UserService
	organizationService  services.OrganizationService
	projectService       services.ProjectService
	fieldService         services.FieldService
	visualizationService services.VisualizationService
	valueService         services.ValueService
	logService           services.LogService
}

func InitServiceProvider(dialector gorm.Dialector) services.ServiceProvider {
	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return nil
	}

	if err := database.SetupJoinTable(&models.User{}, "Projects", &models.UserProject{}); err != nil {
		fmt.Println("Error setting up user projects:", err)
	}

	if err := database.SetupJoinTable(&models.User{}, "Organizations", &models.UserOrganization{}); err != nil {
		fmt.Println("Error setting up user organizations:", err)
	}

	if err := database.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("Error migrating users:", err)
	}

	if err := database.AutoMigrate(&models.Session{}); err != nil {
		fmt.Println("Error migrating sessions:", err)
	}

	if err := database.AutoMigrate(&models.Organization{}); err != nil {
		fmt.Println("Error migrating organizations:", err)
	}

	if err := database.AutoMigrate(&models.Project{}); err != nil {
		fmt.Println("Error migrating projects:", err)
	}

	if err := database.AutoMigrate(&models.Field{}); err != nil {
		fmt.Println("Error migrating fields:", err)
	}

	if err := database.AutoMigrate(&models.Value{}); err != nil {
		fmt.Println("Error migrating values:", err)
	}

	if err := database.AutoMigrate(&models.Visualization{}); err != nil {
		fmt.Println("Error migrating visualizations:", err)
	}

	if err := database.AutoMigrate(&models.Log{}); err != nil {
		fmt.Println("Error migrating logs:", err)
	}

	serviceProvider := &ServiceProvider{}
	serviceProvider.sessionService = &db.SessionService{DB: database}
	serviceProvider.userService = &db.UserService{DB: database}
	serviceProvider.organizationService = &db.OrganizationService{DB: database}
	serviceProvider.projectService = &db.ProjectService{DB: database}
	serviceProvider.fieldService = &db.FieldService{DB: database}
	serviceProvider.valueService = &db.ValueService{DB: database}
	serviceProvider.visualizationService = &db.VisualizationService{DB: database}
	serviceProvider.logService = &db.LogService{DB: database}

	return serviceProvider
}

func (serviceProvider *ServiceProvider) GetSessionService() services.SessionService {
	return serviceProvider.sessionService
}

func (serviceProvider *ServiceProvider) GetUserService() services.UserService {
	return serviceProvider.userService
}

func (serviceProvider *ServiceProvider) GetOrganizationService() services.OrganizationService {
	return serviceProvider.organizationService
}

func (serviceProvider *ServiceProvider) GetProjectService() services.ProjectService {
	return serviceProvider.projectService
}

func (serviceProvider *ServiceProvider) GetFieldService() services.FieldService {
	return serviceProvider.fieldService
}

func (serviceProvider *ServiceProvider) GetVisualizationService() services.VisualizationService {
	return serviceProvider.visualizationService
}

func (serviceProvider *ServiceProvider) GetValueService() services.ValueService {
	return serviceProvider.valueService
}

func (serviceProvider *ServiceProvider) GetLogService() services.LogService {
	return serviceProvider.logService
}
