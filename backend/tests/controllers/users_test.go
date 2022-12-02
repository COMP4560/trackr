package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/tests"
)

func TestGetUserRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "GET", "/api/users/"

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

	response, _ = json.Marshal(responses.User{
		Email:     suite.User.Email,
		FirstName: suite.User.FirstName,
		LastName:  suite.User.LastName,
		CreatedAt: suite.User.CreatedAt,

		NumberOfFields: 1,
		NumberOfValues: 1,

		MaxValueInterval: suite.User.MaxValueInterval,
		MaxValues:        suite.User.MaxValues,
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())
}

func TestUpdateUserRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "PUT", "/api/users/"

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
	// Test invalid parameters path.
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
	// Test empty first name path.
	//

	request, _ := json.Marshal(requests.UpdateUser{})
	response, _ = json.Marshal(responses.Error{Error: "Your first name cannot be empty."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err := suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, suite.User.FirstName, user.FirstName)
	assert.Equal(t, suite.User.LastName, user.LastName)
	assert.Equal(t, suite.User.Password, user.Password)

	//
	// Test empty last name path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		FirstName: suite.User.FirstName,
	})
	response, _ = json.Marshal(responses.Error{Error: "Your last name cannot be empty."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, suite.User.FirstName, user.FirstName)
	assert.Equal(t, suite.User.LastName, user.LastName)
	assert.Equal(t, suite.User.Password, user.Password)

	//
	// Test updating firstname path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		FirstName: "NewFirstName",
		LastName:  suite.User.LastName,
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "NewFirstName", user.FirstName)
	assert.Equal(t, suite.User.LastName, user.LastName)
	assert.Equal(t, suite.User.Password, user.Password)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, logs[0].ID, suite.Logs[0].ID+1)
	assert.Equal(t, "Updated user settings.", logs[0].Message)

	//
	// Test updating lastname path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		LastName:  "NewLastName",
		FirstName: user.FirstName,
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "NewFirstName", user.FirstName)
	assert.Equal(t, "NewLastName", user.LastName)
	assert.Equal(t, suite.User.Password, user.Password)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, logs[0].ID, suite.Logs[0].ID+2)
	assert.Equal(t, "Updated user settings.", logs[0].Message)

	//
	// Test no current password with new password path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		NewPassword: "Password2",
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "NewFirstName", user.FirstName)
	assert.Equal(t, "NewLastName", user.LastName)
	assert.Equal(t, suite.User.Password, user.Password)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, logs[0].ID, suite.Logs[0].ID+3)
	assert.Equal(t, "Updated user settings.", logs[0].Message)

	//
	// Test current password with no new password path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		CurrentPassword: "Password2",
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "NewFirstName", user.FirstName)
	assert.Equal(t, "NewLastName", user.LastName)
	assert.Equal(t, suite.User.Password, user.Password)

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, logs[0].ID, suite.Logs[0].ID+4)
	assert.Equal(t, "Updated user settings.", logs[0].Message)

	//
	// Test updating new password with an invalid current password path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		CurrentPassword: "InvalidPassword",
		NewPassword:     "Password2",
	})
	response, _ = json.Marshal(responses.Error{Error: "Incorrect current password."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful updating password path.
	//

	request, _ = json.Marshal(requests.UpdateUser{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		CurrentPassword: "Password",
		NewPassword:     "Password2",
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "NewFirstName", user.FirstName)
	assert.Equal(t, "NewLastName", user.LastName)
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("Password2")))

	logs, err = suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, logs[0].ID, suite.Logs[0].ID+5)
	assert.Equal(t, "Updated user settings.", logs[0].Message)
}

func TestDeleteUserRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "DELETE", "/api/users/"

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

	user, err := suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, suite.User.ID, user.ID)

	projects, err := suite.Service.GetProjectService().GetProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(projects))

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err = suite.Service.GetUserService().GetUser(suite.User.Email)
	assert.NotNil(t, err)
	assert.Nil(t, user)

	projects, err = suite.Service.GetProjectService().GetProjects(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(projects))
}
