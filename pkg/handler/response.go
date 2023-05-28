package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println("error response:", message)
	c.JSON(statusCode, map[string]interface{}{
		"error": message,
	})
}
