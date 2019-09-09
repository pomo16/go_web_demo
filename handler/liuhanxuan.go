package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	fmt.Println("liuhanxuan")
	c.JSON(200, gin.H{
		"message": "Test success!",
	})
}
