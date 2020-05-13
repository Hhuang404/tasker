package constant

import "time"

const (
	QuerySuccess = "查询成功"
	AddSuccess   = "添加成功"

	ParamError      = "参数异常"
	ParamNull       = "参数为空"
	DuplicateEntity = "重复数据"
	Unauthorized    = "权限不足"

	DoMain = "localhost"

	MaxAge       = 7 * 24 * time.Hour // 七天
	MaxAgeSecond = 604800             // 七天 秒

	YearMonthDayHourMinuteSecond = "2006-01-02 15:04:05"
)
