package handler

import (
	"code.byted.org/gopkg/logs"
	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context) {
	logs.CtxInfo(c, "a sample app log")
	c.JSON(200, gin.H{
		"message": "Get success!",
	})
}
