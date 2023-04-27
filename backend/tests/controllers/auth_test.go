package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"trackr/src/forms/requests"
	"trackr/src/forms/responses"
	"trackr/tests"
)

func TestIsLoggedInRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "GET", "/api/auth/"

	//
	// Test logged in path.
	//

	response, _ := json.Marshal(responses.Empty{})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test not logged in path.
	//

	response, _ = json.Marshal(responses.Error{Error: "Not logged in."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())
}

func TestRegisterRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "POST", "/api/auth/register"

	//
	// Test already logged in path.
	//

	response, _ := json.Marshal(responses.Error{Error: "You cannot make a new account while you are currently logged in."})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no request body path.
	//

	response, _ = json.Marshal(responses.Error{Error: "Invalid request parameters provided."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no first name path.
	//

	request, _ := json.Marshal(requests.Register{})
	response, _ = json.Marshal(responses.Error{Error: "You must provide a first name."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no last name path.
	//

	request, _ = json.Marshal(requests.Register{FirstName: "TestFirstName"})
	response, _ = json.Marshal(responses.Error{Error: "You must provide a last name."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no email path.
	//

	request, _ = json.Marshal(requests.Register{FirstName: "FirstName", LastName: "LastName"})
	response, _ = json.Marshal(responses.Error{Error: "You must provide an email address."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no email path.
	//

	request, _ = json.Marshal(requests.Register{FirstName: "FirstName", LastName: "LastName"})
	response, _ = json.Marshal(responses.Error{Error: "You must provide an email address."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test no password path.
	//

	request, _ = json.Marshal(requests.Register{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "Email",
	})
	response, _ = json.Marshal(responses.Error{Error: "You must provide a password."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid email path.
	//

	request, _ = json.Marshal(requests.Register{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "Email",
		Password:  "Password",
	})
	response, _ = json.Marshal(responses.Error{Error: "You must provide a valid email address."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test an already taken email address
	//

	request, _ = json.Marshal(requests.Register{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     suite.User.Email,
		Password:  "Password",
	})
	response, _ = json.Marshal(responses.Error{Error: "The email is already taken."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful path.
	//

	request, _ = json.Marshal(requests.Register{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "email@email.com",
		Password:  "Password",
	})
	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	user, err := suite.Service.GetUserService().GetUser("email@email.com")
	assert.NotNil(t, user)
	assert.Nil(t, err)

	assert.Equal(t, user.ID, uint(2))
	assert.Equal(t, user.Email, "email@email.com")
	assert.Equal(t, user.FirstName, "FirstName")
	assert.Equal(t, user.LastName, "LastName")
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("Password")))

	logs, err := suite.Service.GetLogService().GetLogs(*user)
	assert.Nil(t, err)
	assert.Equal(t, "Created a new account.", logs[0].Message)
}

func TestLogoutRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "GET", "/api/auth/logout"

	//
	// Test no cookie path.
	//

	response, _ := json.Marshal(responses.Error{Error: "No session cookie provided."})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test not logged in path.
	//

	response, _ = json.Marshal(responses.Error{Error: "You are currently not logged in."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=InvalidSessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test not logged in path.
	//

	response, _ = json.Marshal(responses.Error{Error: "You are currently not logged in."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=InvalidSessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test successful logout path.
	//

	session, _, err := suite.Service.GetSessionService().GetSessionAndUser("ExpiredSessionID")
	assert.Nil(t, err)
	assert.NotNil(t, session)

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser("SessionID")
	assert.Nil(t, err)
	assert.NotNil(t, session)

	response, _ = json.Marshal(responses.Empty{})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser("ExpiredSessionID")
	assert.NotNil(t, err)
	assert.Nil(t, session)

	session, _, err = suite.Service.GetSessionService().GetSessionAndUser("SessionID")
	assert.NotNil(t, err)
	assert.Nil(t, session)

	logs, err := suite.Service.GetLogService().GetLogs(suite.User)
	assert.Nil(t, err)
	assert.Equal(t, "Signed out.", logs[0].Message)
}

func TestLoginRoute(t *testing.T) {
	suite := tests.StartupWithRouter(t)
	method, path := "POST", "/api/auth/login"

	//
	// Test already logged in path.
	//

	response, _ := json.Marshal(responses.Error{Error: "You are currently logged in."})
	httpRecorder := httptest.NewRecorder()
	httpRequest, _ := http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=SessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusTemporaryRedirect, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test expired session cookie path.
	//

	response, _ = json.Marshal(responses.Error{Error: "Invalid request parameters provided."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	httpRequest.Header.Add("Cookie", "Session=ExpiredSessionID")
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test invalid request parameters path.
	//

	response, _ = json.Marshal(responses.Error{Error: "Invalid request parameters provided."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, nil)
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test not passing an email path.
	//

	request, _ := json.Marshal(requests.Login{Email: "", Password: "", RememberMe: false})
	response, _ = json.Marshal(responses.Error{Error: "You must provide an email."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test not passing a password path.
	//

	request, _ = json.Marshal(requests.Login{Email: suite.User.Email, Password: "", RememberMe: false})
	response, _ = json.Marshal(responses.Error{Error: "You must provide a password."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusBadRequest, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test passing a valid email but an invalid password.
	//

	request, _ = json.Marshal(requests.Login{Email: suite.User.Email, Password: "InvalidPassword", RememberMe: false})
	response, _ = json.Marshal(responses.Error{Error: "Invalid email or password combination."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test passing a valid password but an invalid email.
	//

	request, _ = json.Marshal(requests.Login{Email: "InvalidEmail", Password: "Password", RememberMe: false})
	response, _ = json.Marshal(responses.Error{Error: "Invalid email or password combination."})
	httpRecorder = httptest.NewRecorder()
	httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
	suite.Router.ServeHTTP(httpRecorder, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpRecorder.Code)
	assert.Equal(t, response, httpRecorder.Body.Bytes())

	//
	// Test if an expired session exists, this will prove that
	// the next test removes expired sessions.
	//

	_, _, err := suite.Service.GetSessionService().GetSessionAndUser("ExpiredSessionID")
	assert.Nil(t, err)

	//
	// Test passing a valid username, password, and rememberMe path.
	//

	for _, shouldRememberMe := range []bool{false, true} {
		request, _ = json.Marshal(requests.Login{
			Email:      suite.User.Email,
			Password:   "Password",
			RememberMe: shouldRememberMe,
		})
		response, _ = json.Marshal(responses.Empty{})
		httpRecorder = httptest.NewRecorder()
		httpRequest, _ = http.NewRequest(method, path, bytes.NewReader(request))
		suite.Router.ServeHTTP(httpRecorder, httpRequest)

		setCookie := httpRecorder.Header().Get("Set-Cookie")
		sessionId := setCookie[strings.Index(setCookie, "=")+1 : strings.Index(setCookie, ";")]

		assert.Equal(t, http.StatusOK, httpRecorder.Code)
		assert.Equal(t, response, httpRecorder.Body.Bytes())

		session, user, err := suite.Service.GetSessionService().GetSessionAndUser(sessionId)
		assert.Equal(t, user.ID, suite.User.ID)
		assert.Equal(t, session.ID, sessionId)

		if shouldRememberMe {
			assert.GreaterOrEqual(t, session.ExpiresAt.Add(time.Millisecond*time.Duration(100)), session.CreatedAt.AddDate(0, 1, 0))
		} else {
			assert.LessOrEqual(t, session.ExpiresAt, session.CreatedAt.AddDate(0, 0, 7))
			assert.GreaterOrEqual(t, session.ExpiresAt.Add(time.Millisecond*time.Duration(100)), session.CreatedAt.AddDate(0, 0, 7))
		}

		assert.Nil(t, err)

		_, _, err = suite.Service.GetSessionService().GetSessionAndUser("ExpiredSessionID")
		assert.NotNil(t, err)

		logs, err := suite.Service.GetLogService().GetLogs(suite.User)
		assert.Nil(t, err)
		assert.Equal(t, "Signed in.", logs[0].Message)
	}
}
