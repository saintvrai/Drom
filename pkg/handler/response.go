package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/saintvrai/Drom/pkg/logging"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log := logging.GetLogger()
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
