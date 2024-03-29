package dao

import (
	"lhx-github/go_web_demo/model"
)

// 统一接口使用数组类型返回
func GetUserInfoById(userId string) ([]*model.UserInfo, error) {
	var userInfo model.UserInfo
	var userInfos []*model.UserInfo
	err := db.Table("user_info").Where("id = ?", userId).Find(&userInfo).Error
	if err != nil {
		return nil, err
	} else {
		return append(userInfos, &userInfo), nil
	}
}

func GetUserInfoByName(userName string) ([]*model.UserInfo, error) {
	var userInfos []*model.UserInfo
	rows, err := db.Raw("SELECT * FROM goweb.user_info WHERE name = ?", userName).Rows()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userInfo model.UserInfo
		db.ScanRows(rows, &userInfo)
		userInfos = append(userInfos, &userInfo)
	}
	return userInfos, nil
}

func GetUserInfoByAge(userAge int32) ([]*model.UserInfo, error) {
	var userInfos []*model.UserInfo
	rows, err := db.Raw("SELECT * FROM goweb.user_info WHERE age = ?", userAge).Rows()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userInfo model.UserInfo
		db.ScanRows(rows, &userInfo)
		userInfos = append(userInfos, &userInfo)
	}
	return userInfos, nil
}
