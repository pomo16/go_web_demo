package handler

import (
	"github.com/gin-gonic/gin"
	"lhx-github/go_web_demo/consts"
	"lhx-github/go_web_demo/exceptions"
	"lhx-github/go_web_demo/executor"
	"lhx-github/go_web_demo/model"
	"lhx-github/go_web_demo/processor"
)

func UserInfoList(c *gin.Context) {

	context := model.NewUserInfoListContext()

	preContextLoader := &processor.PreContextLoader{}
	preContextCode := executor.LoaderCommon(c, context, preContextLoader)
	if preContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(preContextCode)
		c.JSON(200, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	userInfoListContextLoader := &processor.UserInfoListLoader{}
	userInfoListContextCode := executor.LoaderCommon(c, context, userInfoListContextLoader)
	if userInfoListContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(userInfoListContextCode)
		c.JSON(200, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  consts.MsgSuccess,
		"data":     packUserInfoList(c, context),
		"err_no":   0,
		"err_tips": "成功",
	})
}

func packUserInfoList(c *gin.Context, context model.IUserInfoListContext) map[string]interface{} {
	userInfoList := context.GetUserInfoList()
	listMap := make([]map[string]interface{}, len(userInfoList))
	for key, val := range userInfoList {
		listMap[key] = map[string]interface{}{
			"user_id":   val.UserId,
			"user_name": val.UserName,
			"user_age":  val.Age,
		}
	}

	result := map[string]interface{}{
		"userinfo_list": listMap,
	}

	return result
}
