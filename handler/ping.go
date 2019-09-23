package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Ping is a test interface
func Ping(c *gin.Context) {
	logrus.Warn("ping")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}