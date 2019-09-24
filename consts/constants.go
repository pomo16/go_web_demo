package consts

//消息响应
const (
	MsgSuccess = "success"
	MsgError   = "error"
)

//QueryType 反馈查询类型
type QueryType int16

//QueryType 查询类型
const (
	IdType   = 1 //按主键单条查找
	NameType = 2 //按姓名单条查找
	AgeType  = 3 //按年龄多条查询
)
