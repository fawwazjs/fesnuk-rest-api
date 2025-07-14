package handlers

import (
	"github.com/gin-gonic/gin"
)

func sendResponse(c *gin.Context, httpStatus int, message string, data interface{}, err interface{}) {
	c.JSON(httpStatus, gin.H{
		"message": message,
		"data":    data,
		"error":   err,
	})
}