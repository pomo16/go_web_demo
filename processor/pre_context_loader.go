package processor

import (
	"github.com/gin-gonic/gin"
	"lhx-github/go_web_demo/exceptions"
	"lhx-github/go_web_demo/model"
	"lhx-github/go_web_demo/service"
)

//PreContextLoader 请求预处理
type PreContextLoader struct{}

//Process 预处理过程
func (dif *PreContextLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	parameter := service.ParseInputParameter(ctx)
	runCtx.SetInputParameter(parameter)
	return nil
}

//Name Processor名称
func (dif *PreContextLoader) Name() string {
	return "PreContextLoader"
}
