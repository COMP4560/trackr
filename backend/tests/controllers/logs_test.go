package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"trackr/src/forms/responses"
	"trackr/tests"
)

func TestGetLogsRoute(t *testing.T) {
	suite := tests.StartupWithRouter()
	method, path := "GET", "/api/logs/"

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

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, len(suite.Logs), len(logs))

	response, _ = json.Marshal(responses.LogList{
		Logs: []responses.Log{
			{
				Message:   logs[0].Message,
				CreatedAt: logs[0].CreatedAt,

				ProjectID:   logs[0].ProjectID,
				ProjectName: suite.Project.Name,
			},
			{
				Message:   logs[1].Message,
				CreatedAt: logs[1].CreatedAt,

				ProjectID:   nil,
				ProjectName: "",
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
