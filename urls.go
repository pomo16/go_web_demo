package main

import (
	"github.com/gin-gonic/gin"
	"lhx-github/go_web_demo/handler"
)

var urlMaps = map[string]gin.HandlerFunc{
	"/ping": handler.Ping,
	"/usl":  handler.UserInfoList,
}

func InstanceRoutine() *gin.Engine {
	r := gin.New()
	for url, handler := range urlMaps {
		r.GET(url, handler)
		r.POST(url, handler)
	}
	return r
}
