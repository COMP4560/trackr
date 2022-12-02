package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/models"
	"trackr/tests"
)

func TestAddVisualizationRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "POST", "/api/visualizations/"

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
	// Test invalid request parameters.
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
	// Test empty metadata parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "The metadata parameter cannot be empty.",
	})
	request, _ := json.Marshal(requests.AddVisualization{})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid field id parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find field.",
	})
	request, _ = json.Marshal(requests.AddVisualization{
		FieldID:  models.Field{}.ID,
		Metadata: "Metadata",
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

	visualization, err := suite.Service.GetVisualizationService().GetVisualization(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, visualization)

	response, _ = json.Marshal(responses.NewVisualization{
		ID: uint(2),
	})
	request, _ = json.Marshal(requests.AddVisualization{
		FieldID:  suite.Field.ID,
		Metadata: "Metadata",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(2, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)
	assert.Equal(t, uint(2), visualization.ID)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Added a new visualization.", logs[0].Message)
}

func TestGetVisualizationsRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "GET", "/api/visualizations/"

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
	// Test invalid project id parameter.
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

	newVisualization := suite.Visualization
	newVisualization.ID = 2

	visualizationId, err := suite.Service.GetVisualizationService().AddVisualization(newVisualization)
	assert.Nil(t, err)
	assert.Equal(t, newVisualization.ID, visualizationId)

	response, _ = json.Marshal(responses.VisualizationList{
		Visualizations: []responses.Visualization{
			{
				ID:        suite.Visualization.ID,
				FieldID:   suite.Field.ID,
				Metadata:  suite.Visualization.Metadata,
				UpdatedAt: suite.Visualization.UpdatedAt,
				CreatedAt: suite.Visualization.CreatedAt,
			},
			{
				ID:        newVisualization.ID,
				FieldID:   suite.Field.ID,
				Metadata:  newVisualization.Metadata,
				UpdatedAt: newVisualization.UpdatedAt,
				CreatedAt: newVisualization.CreatedAt,
			},
		},
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+fmt.Sprint(suite.Project.ID), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())
}

func TestUpdateVisualizationsRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "PUT", "/api/visualizations/"

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
	// Test empty metadata path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "The metadata parameter cannot be empty.",
	})
	request, _ := json.Marshal(requests.UpdateVisualization{})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non-existant visulization id path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find visualization.",
	})
	request, _ = json.Marshal(requests.UpdateVisualization{
		ID:       models.Visualization{}.ID,
		Metadata: "NewMetadata",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non-existant field id path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find corresponding field.",
	})
	request, _ = json.Marshal(requests.UpdateVisualization{
		ID:       suite.Visualization.ID,
		Metadata: "NewMetadata",
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

	newField := suite.Field
	newField.ID = 2

	newFieldId, err := suite.Service.GetFieldService().AddField(newField)
	assert.Nil(t, err)
	assert.Equal(t, newField.ID, newFieldId)

	visualization, err := suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)
	assert.Equal(t, suite.Visualization.ID, visualization.ID)
	assert.Equal(t, suite.Visualization.Metadata, visualization.Metadata)

	response, _ = json.Marshal(responses.Empty{})
	request, _ = json.Marshal(requests.UpdateVisualization{
		ID:       suite.Visualization.ID,
		FieldID:  newFieldId,
		Metadata: "NewMetadata",
	})

	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)
	assert.Equal(t, suite.Visualization.ID, visualization.ID)
	assert.Equal(t, newFieldId, visualization.FieldID)
	assert.Equal(t, "NewMetadata", visualization.Metadata)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Modified a visualization.", logs[0].Message)
}

func TestDeleteVisualizationRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "DELETE", "/api/visualizations/"

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
	// Test invalid visualization id parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Invalid :visualizationId parameter provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"invalid", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non-existant visualization id path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find visualization.",
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

	visualization, err := suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, visualization)
	assert.Equal(t, suite.Visualization.ID, visualization.ID)

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+fmt.Sprint(suite.Visualization.ID), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	visualization, err = suite.Service.GetVisualizationService().GetVisualization(suite.Visualization.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, visualization)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Deleted a visualization.", logs[0].Message)
}
