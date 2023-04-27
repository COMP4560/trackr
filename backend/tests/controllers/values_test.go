package controllers_test

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"time"
	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/src/models"
	"trackr/tests"
)

func TestGetValuesRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "GET", "/api/values/"

	//
	// Test invalid order parameter path.
	//

	request, _ := query.Values(requests.GetValues{
		Order: "invalid",
	})
	response, _ := json.Marshal(responses.Error{
		Error: "Invalid order parameter provided.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path+"?"+request.Encode(), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid offset parameter path.
	//

	request, _ = query.Values(requests.GetValues{
		Order:  "asc",
		Offset: -1,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Invalid offset parameter provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"?"+request.Encode(), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid limit parameter path.
	//

	request, _ = query.Values(requests.GetValues{
		Order:  "asc",
		Offset: 0,
		Limit:  -1,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Invalid limit parameter provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"?"+request.Encode(), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid project api parameter path.
	//

	request, _ = query.Values(requests.GetValues{
		Order:   "asc",
		Offset:  0,
		Limit:   0,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find project, invalid API key.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"?"+request.Encode(), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid field id parameter path.
	//

	request, _ = query.Values(requests.GetValues{
		APIKey:  suite.UserProject.APIKey,
		Order:   "asc",
		Offset:  0,
		Limit:   0,
		FieldID: models.Field{}.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find field.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"?"+request.Encode(), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	expectedValues := []models.Value{suite.Value}
	numberOfExpectedValues := 10 + len(expectedValues)

	for i := 2; i <= numberOfExpectedValues; i++ {
		newValue := suite.Value
		newValue.ID = uint(i)
		newValue.CreatedAt = time.Now()
		newValue.Value = fmt.Sprintf("%d.00", i)

		expectedValues = append(expectedValues, newValue)

		suite.Service.GetValueService().AddValue(newValue)
		time.Sleep(time.Second)
	}

	for offset := 0; offset < numberOfExpectedValues; offset++ {
		for limit := 0; limit <= numberOfExpectedValues; limit++ {
			for _, order := range []string{"desc", "asc"} {
				expectedValuesSlice := expectedValues[:]

				sort.Slice(expectedValuesSlice, func(i, j int) bool {
					if order == "desc" {
						return expectedValuesSlice[i].CreatedAt.After(
							expectedValuesSlice[j].CreatedAt,
						)
					} else {
						return expectedValuesSlice[i].CreatedAt.Before(
							expectedValuesSlice[j].CreatedAt,
						)
					}
				})

				if offset == 0 {
					expectedValuesSlice = expectedValuesSlice[:]
				} else {
					expectedValuesSlice = expectedValuesSlice[offset:]
				}

				request, _ = query.Values(requests.GetValues{
					APIKey:  suite.UserProject.APIKey,
					Order:   order,
					Offset:  offset,
					Limit:   limit,
					FieldID: suite.Field.ID,
				})

				httpRecorder = httptest.NewRecorder()
				httpRequest, _ = http.NewRequest(method, path+"?"+request.Encode(), nil)
				httpRequest.Header.Add("Cookie", "Session=SessionID")
				suite.Router.ServeHTTP(httpRecorder, httpRequest)

				assert.Equal(t, http.StatusOK, httpRecorder.Code)

				var response responses.ValueList
				err := json.Unmarshal(httpRecorder.Body.Bytes(), &response)
				assert.Nil(t, err)

				if limit > 0 {
					assert.LessOrEqual(t, len(response.Values), limit)
				} else {
					assert.Equal(t, len(expectedValuesSlice), len(response.Values))
				}

				for i := 0; i < len(response.Values); i++ {
					assert.Equal(t, expectedValuesSlice[i].Value, response.Values[i].Value)
				}
			}
		}
	}
}

func TestDeleteValuesRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "DELETE", "/api/values/"

	//
	// Test not logged in path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Not authorized to access this resource.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path+"invalid", nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid field id parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Invalid :fieldId parameter provided.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+"invalid", nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test non existant field parameter path.
	//

	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find field.",
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

	value, err := suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, suite.Value.ID, value.ID)

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path+fmt.Sprintf("%d", suite.Field.ID), nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	value, err = suite.Service.GetValueService().GetValue(suite.Value.ID, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"Delete all values associated to the field %s.", suite.Field.Name), logs[0].Message)
}

func TestAddValueRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "POST", "/api/values/"

	//
	// Test invalid request parameters path.
	//

	response, _ := json.Marshal(responses.Error{
		Error: "Invalid request parameters provided.",
	})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test empty value path.
	//

	request, _ := query.Values(requests.AddValue{})
	response, _ = json.Marshal(responses.Error{
		Error: "The value cannot be empty.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid api key path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:  "2.00",
		APIKey: "invalid",
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find project, invalid API key.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid field id parameter path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:   "2.00",
		APIKey:  suite.UserProject.APIKey,
		FieldID: models.Field{}.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "Failed to find field.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid value type path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:   "invalid",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "The value must be a floating point number.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test NaN value path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:   "NaN",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "The value cannot be NaN nor Infinity.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test -Inf value path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:   "-Inf",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "The value cannot be NaN nor Infinity.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test Inf value path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:   "Inf",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "The value cannot be NaN nor Infinity.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test maximum values path.
	//

	request, _ = query.Values(requests.AddValue{
		Value:   "2.00",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Error{
		Error: "You have exceeded your max values limit.",
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	suite.User.MaxValues++
	suite.Service.GetUserService().UpdateUser(suite.User)

	value, err := suite.Service.GetValueService().GetValue(2, suite.User)
	assert.NotNil(t, err)
	assert.Nil(t, value)

	request, _ = query.Values(requests.AddValue{
		Value:   "2.00",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	value, err = suite.Service.GetValueService().GetValue(2, suite.User)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, uint(2), value.ID)
	assert.Equal(t, "2.00", value.Value)

	//
	// Test rate limit.
	//

	suite.User.MaxValues++
	suite.Service.GetUserService().UpdateUser(suite.User)

	request, _ = query.Values(requests.AddValue{
		Value:   "3.00",
		APIKey:  suite.UserProject.APIKey,
		FieldID: suite.Field.ID,
	})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, strings.NewReader(request.Encode()))
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusTooManyRequests, httpRecorder.Code)
}
