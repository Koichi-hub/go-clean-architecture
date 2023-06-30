package middlewares_tests

import (
	"bytes"
	"go-clean-architecture/controllers/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	type args struct {
		sessionId string
	}

	type testCase struct {
		name string
		args
		expected string
	}

	tests := []testCase{
		{
			name: "Getting session id",
			args: args{
				sessionId: "1",
			},
			expected: "1",
		},
	}

	gin.SetMode(gin.ReleaseMode)
	sessionMiddleware := middlewares.NewSessionMiddleware()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

			req, err := http.NewRequest("GET", "/tasks", bytes.NewBuffer([]byte{}))
			assert.NoError(t, err)
			req.AddCookie(&http.Cookie{
				Name:   "session_id",
				Value:  test.args.sessionId,
				MaxAge: 300,
			})

			ctx.Request = req

			sessionMiddleware.Middleware()(ctx)
			sessionId, exists := ctx.Get("sessionId")

			assert.True(t, exists)
			assert.EqualValues(t, test.expected, sessionId)
		})
	}
}
