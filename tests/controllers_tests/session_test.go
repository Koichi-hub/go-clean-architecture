package controllers_tests

import (
	"bytes"
	"go-clean-architecture/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	sessionController := controllers.NewSessionController()

	res := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(res)

	req, err := http.NewRequest("POST", "/sessions/login", bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	ctx.Request = req

	sessionController.Login(ctx)
	assert.Equal(t, http.StatusOK, res.Code)

	var receivedSessionId bool
	cookies := res.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "session_id" && cookie.Value != "" {
			receivedSessionId = true
			break
		}
	}
	assert.True(t, receivedSessionId)
}

func TestLogout(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	sessionController := controllers.NewSessionController()

	res := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(res)

	req, err := http.NewRequest("POST", "/sessions/logout", bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	ctx.Request = req

	sessionController.Logout(ctx)
	assert.Equal(t, http.StatusOK, res.Code)

	var receivedEmptySessionId bool
	cookies := res.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "session_id" && cookie.Value == "" && cookie.MaxAge == -1 {
			receivedEmptySessionId = true
			break
		}
	}
	assert.True(t, receivedEmptySessionId)
}
