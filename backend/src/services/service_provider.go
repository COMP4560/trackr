package services

type ServiceProvider interface {
	GetSessionService() SessionService
	GetProjectService() ProjectService
	GetUserService() UserService
	GetFieldService() FieldService
	GetLogService() LogService
	GetVisualizationService() VisualizationService
	GetValueService() ValueService
}
