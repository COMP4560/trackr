package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/tests"
)

func TestAddFieldRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "POST", "/api/fields/"

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

	field, err := suite.Service.GetFieldService().GetField(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, field)

	//
	// Test invalid project id paramater.
	//

	request, _ := json.Marshal(requests.AddField{
		ProjectID: 0,
		Name:      "Field2",
	})

	response, _ = json.Marshal(responses.Error{
		Error: "Cannot find project.",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test missing name paramater.
	//

	request, _ = json.Marshal(requests.AddField{
		ProjectID: 1,
		Name:      "",
	})

	response, _ = json.Marshal(responses.Error{
		Error: "The name of a field cannot be empty.",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	request, _ = json.Marshal(requests.AddField{
		ProjectID: 1,
		Name:      "Field2",
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	field, err = suite.Service.GetFieldService().GetField(2, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, field)
	assert.Equal(t, uint(2), field.ID)
}

func TestGetFieldsRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "GET", "/api/fields/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})

	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path+"1", nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
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

	newField := suite.Field
	newField.ID = 2
	newField.Name = "Field2"
	newField.CreatedAt = suite.Time
	newField.UpdatedAt = suite.Time
	newField.Project = suite.Project

	err := suite.Service.GetFieldService().AddField(newField)
	assert.Nil(t, err)

	response, _ = json.Marshal(responses.FieldList{
		Fields: []responses.Field{
			{
				ID:   suite.Field.ID,
				Name: suite.Field.Name,
			},
			{
				ID:   newField.ID,
				Name: newField.Name,
			},
		},
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"1", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())
}

func TestUpdateFieldRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "PUT", "/api/fields/"

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

	request, _ := json.Marshal(requests.UpdateField{
		ID: 0,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find field.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no modification path.
	//

	request, _ = json.Marshal(requests.UpdateField{
		ID: 1,
	})

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	field, err := suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, field)
	assert.Equal(t, suite.Field.Name, field.Name)

	//
	// Test updating name path.
	//

	request, _ = json.Marshal(requests.UpdateProject{
		ID:   1,
		Name: "New Field Name",
	})

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	field, err = suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, field)
	assert.Equal(t, "New Field Name", field.Name)
}

func TestDeleteFieldRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "DELETE", "/api/fields/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path+"1", nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid id parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Invalid :id parameter provided.",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"invalid", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non-existant field id path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to delete field.",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"0", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	field, err := suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, field)
	assert.Equal(t, uint(1), field.ID)
	assert.Equal(t, suite.Project.ID, field.ProjectID)

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"1", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	field, err = suite.Service.GetFieldService().GetField(suite.Field.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, field)
}
