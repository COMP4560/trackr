package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/tests"
)

func TestAddProjectRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
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

	project, err := suite.Service.GetProjectService().GetProject(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, project)

	response, _ = json.Marshal(responses.NewProject{
		ID: uint(2),
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	project, err = suite.Service.GetProjectService().GetProject(2, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, uint(2), project.ID)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Created a new project.", logs[0].Message)
}

func TestGetProjectsRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
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

	newProject := suite.Project
	newProject.ID = 2
	newProject.APIKey = "APIKey2"

	projectId, err := suite.Service.GetProjectService().AddProject(newProject)
	assert.Nil(t, err)
	assert.Equal(t, newProject.ID, projectId)

	response, _ = json.Marshal(responses.ProjectList{
		Projects: []responses.Project{
			{
				ID:             suite.Project.ID,
				Name:           suite.Project.Name,
				Description:    suite.Project.Description,
				APIKey:         suite.Project.APIKey,
				CreatedAt:      suite.Project.CreatedAt,
				UpdatedAt:      suite.Project.UpdatedAt,
				NumberOfFields: 1,
			},
			{
				ID:             newProject.ID,
				Name:           newProject.Name,
				Description:    newProject.Description,
				APIKey:         newProject.APIKey,
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
	suite := tests.StartupWithRouter()
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

	response, _ = json.Marshal(responses.Project{
		ID:             suite.Project.ID,
		Name:           suite.Project.Name,
		Description:    suite.Project.Description,
		APIKey:         suite.Project.APIKey,
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
	suite := tests.StartupWithRouter()
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

	project, err := suite.Service.GetProjectService().GetProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, uint(1), project.ID)
	assert.Equal(t, suite.User.ID, project.UserID)

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"1", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	project, err = suite.Service.GetProjectService().GetProject(1, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, project)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("Deleted the project %s.", suite.Project.Name), logs[0].Message)
}

func TestUpdateProjectRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
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
	response, _ = json.Marshal(responses.UpdateProject{APIKey: suite.Project.APIKey})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	project, err := suite.Service.GetProjectService().GetProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "New Project Name", project.Name)
	assert.Equal(t, suite.Project.Description, project.Description)
	assert.Equal(t, suite.Project.APIKey, project.APIKey)

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
	response, _ = json.Marshal(responses.UpdateProject{APIKey: suite.Project.APIKey})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	project, err = suite.Service.GetProjectService().GetProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "New Project Name", project.Name)
	assert.Equal(t, "New Description Name", project.Description)
	assert.Equal(t, suite.Project.APIKey, project.APIKey)

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
	response, _ = json.Marshal(responses.UpdateProject{APIKey: suite.Project.APIKey})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.NotEqual(t, response, httpRecorder.Body.Bytes())

	project, err = suite.Service.GetProjectService().GetProject(1, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "New Project Name", project.Name)
	assert.Equal(t, "New Description Name", project.Description)
	assert.NotEqual(t, suite.Project.APIKey, project.APIKey)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Updated the project's information.", logs[0].Message)
}
