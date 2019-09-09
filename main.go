package main

import (
	"github.com/gin-gonic/gin"
	"lhx-github/go_web_demo/handler"
)

func main() {

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	router := gin.Default()
	//创建不带中间件的路由：
	//r := gin.New()


	router.GET("/someGet", handler.Getting)
	router.GET("/someTest", handler.Test)

	router.Run(":8080")
}