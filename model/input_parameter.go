package model

import "lhx-github/go_web_demo/consts"

type InputParameter struct {
	Id   string
	Name string
	Age  int32
	QueryType consts.QueryType
}
