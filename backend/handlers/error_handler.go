package handlers

import (
	"net/http"

	"github.com/gabrielg2020/backend/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func handleInternalServerError(c *gin.Context, err error, logMessage string) {
	logger.Errorf(logMessage, err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}
