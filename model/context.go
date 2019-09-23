package model

//BaseContext 基础上下文信息
type BaseContext struct {
	inputParameter *InputParameter
}

//IContext 上下文结构读写接口
type IContext interface {
	SetInputParameter(parameter *InputParameter)
	GetInputParameter() *InputParameter
}

// UserInfoListContext 用户信息列表上下文
type UserInfoListContext struct {
	BaseContext
	userInfoList []*UserInfo
}

// IUserInfoListContext 用户信息列表读写接口
type IUserInfoListContext interface {
	IContext
	SetUserInfoList(userInfoList []*UserInfo)
	GetUserInfoList() []*UserInfo
}

// NewUserInfoListContext context构造函数
func NewUserInfoListContext() *UserInfoListContext {
	return &UserInfoListContext{}
}