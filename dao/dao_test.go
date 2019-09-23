package dao

import (
	"fmt"
	"testing"
)

func TestGetUserInfoById(t *testing.T) {
	InitDb()
	defer db.Close()
	userInfo, _ := GetUserInfoById("1")
	fmt.Println(*userInfo)
}

func TestGetUserInfoByName(t *testing.T) {
	InitDb()
	defer db.Close()
	userInfos, _ := GetUserInfoByName("pomo2")
	for _, v := range userInfos {
		fmt.Println(v)
	}
}

func TestGetUserInfoByAge(t *testing.T) {
	InitDb()
	defer db.Close()
	userInfos, _ := GetUserInfoByAge(1)
	for _, v := range userInfos {
		fmt.Println(v)
	}
}
