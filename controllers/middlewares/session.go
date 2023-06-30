package middlewares

import (
	"go-clean-architecture/controllers/controllers_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionMiddleware struct{}

func NewSessionMiddleware() *SessionMiddleware {
	return &SessionMiddleware{}
}

func (*SessionMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie("session_id")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, controllers_errors.NewHttpError("empty cookie: missing session_id"))
			return
		}
		ctx.Set("sessionId", sessionId)
		ctx.Next()
	}
}
