package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"trackr/src/forms/responses/projects"
	"trackr/tests/helpers"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/tests"
)

func TestAddProjectRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "POST", "/api/projects/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	userProject, err := suite.Service.GetProjectService().GetUserProject(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, userProject)

	response, _ = json.Marshal(projects.NewProject{
		ID: uint(2),
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	userProject, err = suite.Service.GetProjectService().GetUserProject(2, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, userProject)
	assert.NotNil(t, userProject.User)
	assert.NotNil(t, userProject.Project)

	assert.Equal(t, uint(2), userProject.ProjectID)
	assert.Equal(t, uint(1), userProject.UserID)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Created a new project.", logs[0].Message)
}

func TestGetProjectsRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "GET", "/api/projects/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	newProject, newUserProject := helpers.CreateNewProject(*suite)

	projectId, err := suite.Service.GetProjectService().AddProject(newProject, newUserProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	response, _ = json.Marshal(projects.ProjectList{
		Projects: []projects.Project{
			{
				ID:             suite.Project.ID,
				Name:           suite.Project.Name,
				Description:    suite.Project.Description,
				APIKey:         suite.UserProject.APIKey,
				CreatedAt:      suite.Project.CreatedAt,
				UpdatedAt:      suite.Project.UpdatedAt,
				NumberOfFields: 1,
			},
			{
				ID:             newProject.ID,
				Name:           newProject.Name,
				Description:    newProject.Description,
				APIKey:         newUserProject.APIKey,
				CreatedAt:      newProject.CreatedAt,
				UpdatedAt:      newProject.UpdatedAt,
				NumberOfFields: 0,
			},
		},
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())
}

func TestGetProjectRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "GET", "/api/projects/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid project id parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Invalid :projectId parameter provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"invalid", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non-existant project id path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find project.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"0", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	response, _ = json.Marshal(projects.Project{
		ID:             suite.Project.ID,
		Name:           suite.Project.Name,
		Description:    suite.Project.Description,
		APIKey:         suite.UserProject.APIKey,
		CreatedAt:      suite.Project.CreatedAt,
		UpdatedAt:      suite.Project.UpdatedAt,
		NumberOfFields: 1,
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"1", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())
}

func TestDeleteProjectRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "DELETE", "/api/projects/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path+"0", nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid project id parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Invalid :projectId parameter provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"invalid", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non-existent project id path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find project.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"0", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	userProject, err := suite.Service.GetProjectService().GetUserProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, userProject)

	assert.NotNil(t, userProject.User)
	assert.NotNil(t, userProject.Project)

	assert.Equal(t, uint(1), userProject.ProjectID)
	assert.Equal(t, uint(1), userProject.UserID)

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"1", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	userProject, err = suite.Service.GetProjectService().GetUserProject(1, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, userProject)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("Deleted the project %s.", suite.Project.Name), logs[0].Message)
}

func TestUpdateProjectRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "PUT", "/api/projects/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid request parameters path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Invalid request parameters provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid project id path.
	//

	request, _ := json.Marshal(requests.UpdateProject{
		ID: 0,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find project.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test empty name path.
	//

	request, _ = json.Marshal(requests.UpdateProject{
		ID: 1,
	})
	response, _ = json.Marshal(responses.Error{Error: "The project's name cannot be empty."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test updating name path.
	//

	request, _ = json.Marshal(requests.UpdateProject{
		ID:          1,
		Name:        "New Project Name",
		Description: suite.Project.Description,
	})
	response, _ = json.Marshal(projects.UpdateProject{APIKey: suite.UserProject.APIKey})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	userProject, err := suite.Service.GetProjectService().GetUserProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, userProject)
	assert.NotNil(t, userProject.User)
	assert.NotNil(t, userProject.Project)

	project := userProject.Project
	assert.Equal(t, "New Project Name", project.Name)
	assert.Equal(t, suite.Project.Description, project.Description)
	assert.Equal(t, suite.UserProject.APIKey, userProject.APIKey)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Updated the project's information.", logs[0].Message)

	//
	// Test updating description path.
	//

	request, _ = json.Marshal(requests.UpdateProject{
		ID:          1,
		Name:        project.Name,
		Description: "New Description Name",
	})
	response, _ = json.Marshal(projects.UpdateProject{APIKey: suite.UserProject.APIKey})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	userProject, err = suite.Service.GetProjectService().GetUserProject(1, suite.User)
	project = userProject.Project
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "New Project Name", project.Name)
	assert.Equal(t, "New Description Name", project.Description)
	assert.Equal(t, suite.UserProject.APIKey, userProject.APIKey)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Updated the project's information.", logs[0].Message)

	//
	// Test resetting api key path.
	//

	request, _ = json.Marshal(requests.UpdateProject{
		ID:          1,
		Name:        project.Name,
		Description: project.Description,
		ResetAPIKey: true,
	})
	response, _ = json.Marshal(projects.UpdateProject{APIKey: suite.UserProject.APIKey})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.NotEqual(t, response, httpRecorder.Body.Bytes())

	userProject, err = suite.Service.GetProjectService().GetUserProject(1, suite.User)
	project = userProject.Project
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "New Project Name", project.Name)
	assert.Equal(t, "New Description Name", project.Description)
	assert.NotEqual(t, suite.UserProject.APIKey, userProject.APIKey)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Updated the project's information.", logs[0].Message)
}
