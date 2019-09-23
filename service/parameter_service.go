package service

import (
	"github.com/gin-gonic/gin"
	"lhx-github/go_web_demo/consts"
	"lhx-github/go_web_demo/model"
	"lhx-github/go_web_demo/utils"
)

func ParseInputParameter(ctx *gin.Context) *model.InputParameter {
	parameter := &model.InputParameter{
		Id: utils.GetParamString(ctx, "id", ""),
		Name: utils.GetParamString(ctx, "name", ""),
		Age: utils.GetParamInt32(ctx, "age", 0),
		QueryType: consts.QueryType(utils.GetParamInt16(ctx, "query_type", 0)),
	}

	return parameter
}
