package service

import (
	"lhx-github/go_web_demo/consts"
	"lhx-github/go_web_demo/dao"
	"lhx-github/go_web_demo/model"
)

func GetUserInfoList(params *model.UserInfoParams) ([]*model.UserInfo, error) {

	var userInfoList []*model.UserInfo
	var err error
	switch params.QueryType {
	case consts.IdType:
		userInfoList, err = dao.GetUserInfoById(params.UserId)
	case consts.NameType:
		userInfoList, err = dao.GetUserInfoByName(params.UserName)
	case consts.AgeType:
		userInfoList, err = dao.GetUserInfoByAge(params.UserAge)
	}

	if err != nil {
		return nil, err
	}

	return userInfoList, nil
}
