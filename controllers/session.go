package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

type SessionController struct{}

func NewSessionController() *SessionController {
	return &SessionController{}
}

func (sessionController *SessionController) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/sessions")
	g.POST("/login", sessionController.Login)
	g.DELETE("/logout", sessionController.Logout)
}

func (sessionController *SessionController) Login(ctx *gin.Context) {
	sessionId := randstr.String(10)
	ctx.SetCookie("session_id", sessionId, int(time.Hour.Seconds()), "/", "", false, true)
	ctx.Status(http.StatusOK)
}

func (sessionController *SessionController) Logout(ctx *gin.Context) {
	ctx.SetCookie("session_id", "", -1, "/", "", false, true)
	ctx.Status(http.StatusOK)
}
