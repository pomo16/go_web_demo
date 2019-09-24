package model

//SetInputParameter 设置输入参数
func (infoCtx *BaseContext) SetInputParameter(parameter *InputParameter) {
	infoCtx.inputParameter = parameter
}

//GetInputParameter 设置输出参数
func (infoCtx *BaseContext) GetInputParameter() *InputParameter {
	if infoCtx.inputParameter != nil {
		return infoCtx.inputParameter
	}
	return &InputParameter{}
}

//SetUserInfoList 设置反馈列表
func (infoCtx *UserInfoListContext) SetUserInfoList(feedbackList []*UserInfo) {
	infoCtx.userInfoList = feedbackList
}

//GetUserInfoList 获取反馈列表
func (infoCtx *UserInfoListContext) GetUserInfoList() []*UserInfo {
	if infoCtx.userInfoList != nil {
		return infoCtx.userInfoList
	}
	return []*UserInfo{}
}
