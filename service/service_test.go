package service

import (
	"fmt"
	"lhx-github/go_web_demo/consts"
	"lhx-github/go_web_demo/dao"
	"lhx-github/go_web_demo/model"
	"testing"
)

func TestGetUserInfoList(t *testing.T) {
	dao.InitDb()
	params := &model.UserInfoParams{
		UserName: "pomo2",
		QueryType: consts.NameType,
	}

	res, _ := GetUserInfoList(params)
	fmt.Println(len(res))
	for _, v := range res {
		fmt.Println(v)
	}
}
