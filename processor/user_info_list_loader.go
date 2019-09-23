package processor

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lhx-github/go_web_demo/consts"
	"lhx-github/go_web_demo/exceptions"
	"lhx-github/go_web_demo/model"
	"lhx-github/go_web_demo/service"
)

//ListLoader 反馈列表loader
type UserInfoListLoader struct{}

//Process 获取列表
func (loader *UserInfoListLoader) Process(ctx *gin.Context, runCtx model.IContext) exceptions.ErrProcessor {
	listCtx, ok := runCtx.(model.IUserInfoListContext)
	if !ok {
		logrus.Warn(ctx, "user info list loader listCtx error")
		return exceptions.ErrTypeAssert
	}

	inputParameter := runCtx.GetInputParameter()
	switch inputParameter.QueryType {
	case consts.IdType:
		if inputParameter.Id == "" {
			return exceptions.ErrRequestParams
		}
	case consts.NameType:
		if inputParameter.Name == "" {
			return exceptions.ErrRequestParams
		}
	case consts.AgeType:
		if inputParameter.Age == 0 {
			return exceptions.ErrRequestParams
		}
	default:
		return exceptions.ErrUnknownType
	}

	userInfoParams := &model.UserInfoParams{
		UserId:   inputParameter.Id,
		UserName: inputParameter.Name,
		UserAge:  inputParameter.Age,
		QueryType: inputParameter.QueryType,
	}

	userInfoList, err := service.GetUserInfoList(userInfoParams)
	logrus.Error(len(userInfoList))
	if userInfoList != nil {
		listCtx.SetUserInfoList(userInfoList)
		return nil
	}

	if err == nil && len(userInfoList) == 0 {
		return exceptions.ErrResultEmpty
	}

	logrus.Errorf("UserInfoList loader return err: %v", err)
	return exceptions.ErrProcessFailed
}

//Name 获取processor名
func (loader *UserInfoListLoader) Name() string {
	return "UserInfoListLoader"
}