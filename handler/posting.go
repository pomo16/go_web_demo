package handler

import (
	"code.byted.org/gopkg/logs"
	"github.com/gin-gonic/gin"
)

func Posting(c *gin.Context) {
	logs.CtxInfo(c, "a sample app log")
	c.JSON(200, gin.H{
		"message": "Post success!",
	})
}
