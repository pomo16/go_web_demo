package executor

import (
	"github.com/gin-gonic/gin"
	"lhx-github/go_web_demo/exceptions"
	"lhx-github/go_web_demo/model"
)

//IProcessor 处理器接口
type IProcessor interface {
	Process(*gin.Context, model.IContext) exceptions.ErrProcessor
	Name() string
}

//LoaderCommon 加载通用任务处理
func LoaderCommon(ctx *gin.Context, runCtx model.IContext, processor IProcessor) exceptions.ErrProcessor {
	code := exceptions.ErrProcessPanic
	code = processor.Process(ctx, runCtx)
	return code
}
