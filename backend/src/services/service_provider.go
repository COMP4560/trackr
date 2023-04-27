package services

type ServiceProvider interface {
	GetSessionService() SessionService
	GetOrganizationService() OrganizationService
	GetProjectService() ProjectService
	GetUserService() UserService
	GetFieldService() FieldService
	GetLogService() LogService
	GetVisualizationService() VisualizationService
	GetValueService() ValueService
}
