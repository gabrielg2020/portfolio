package handlers

import (
	"net/http"

	"github.com/gabrielg2020/backend/logger"
	"github.com/gin-gonic/gin"
)

func handleInternalServerError(c *gin.Context, err error, logMessage string) {
	logger.Errorf(logMessage, err)
	// This condition is only hit when testing how frontend deals with errors
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Fake Erorr",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
