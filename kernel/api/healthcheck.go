package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "healthy",
		"date":    time.Now().Format(time.RFC3339),
	})
}
