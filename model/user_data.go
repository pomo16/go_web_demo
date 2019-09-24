package model

import "lhx-github/go_web_demo/consts"

// 用户信息查询参数
type UserInfoParams struct {
	UserId    string
	UserName  string
	UserAge   int32
	QueryType consts.QueryType
}

// 用户信息结构体
type UserInfo struct {
	UserId   string `gorm:"column:id"`
	UserName string `gorm:"column:name"`
	Age      int32  `gorm:"column:age"`
}
