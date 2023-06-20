package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getSessionId(ctx *gin.Context) (string, error) {
	if sessionId, exists := ctx.Get("sessionId"); exists {
		return sessionId.(string), nil
	}
	return "", fmt.Errorf("missing sessionId")
}
