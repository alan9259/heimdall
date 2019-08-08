package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"heimdall/internal/platform"
	"heimdall/internal/router/middleware"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSignUpCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"first_name":"alan","email_address":"alan@miu.com","password":"1234567"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/accounts", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.SignUp(c))
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "alan@miu.com", m["email_address"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestLoginSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"email_address":"alan@miu.com","password":"1234567"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/accounts/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.Login(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "alan@miu.com", m["email_address"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestLoginFailed(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"email_address":"test123@miu.com","password":"1234567"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/accounts/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.Login(c))
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func TestGetCurrentAccountSuccess(t *testing.T) {
	tearDown()
	setup()

	jwtMiddleware := middleware.JWT(platform.JWTSecret, h.revokedTokenStore)
	req := httptest.NewRequest(echo.GET, "/api/accounts/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(platform.GenerateJWTToken(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.GetCurrentAccount(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "alan@miu.com", m["email_address"])
	}
}

func TestCurrentAccountCaseInvalid(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(platform.JWTSecret, h.revokedTokenStore)
	req := httptest.NewRequest(echo.GET, "/api/accounts/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(platform.GenerateJWTToken(100)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.GetCurrentAccount(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateAccountMultipleFields(t *testing.T) {
	tearDown()
	setup()
	var (
		account1UpdateReq = `{
			"first_name":"alan", 
			"last_name": "wang", 
			"gender_id":2, 
			"phone_number": 111, 
			"date_of_birth": "1985-04-28T10:00:00Z"
		}`
	)

	jwtMiddleware := middleware.JWT(platform.JWTSecret, h.revokedTokenStore)
	req := httptest.NewRequest(echo.PUT, "/api/profiles", strings.NewReader(account1UpdateReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(platform.GenerateJWTToken(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.UpdateAccount(c)
	})(c)
	assert.NoError(t, err)

	err = jwtMiddleware(func(context echo.Context) error {
		return h.GetCurrentAccount(c)
	})(c)

	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "alan", m["first_name"])
		assert.Equal(t, "wang", m["last_name"])
		assert.Equal(t, "male", m["gender"])
		assert.NotEmpty(t, m["token"])
	}
}
