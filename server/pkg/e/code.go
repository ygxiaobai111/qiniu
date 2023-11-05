package e

/*
错误定义
*/
const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400
	ERRORNULL             = 410

	//成员错误
	ErrorExistNick          = 10001
	ErrorExistUser          = 10002
	ErrorNotExistUser       = 10003
	ErrorNotCompare         = 10004
	ErrorNotComparePassword = 10005
	ErrorFailEncryption     = 10006
	ErrorUserNotFound       = 10010

	//管理员错误
	ErrorAdminFindUser = 30010
	//数据库错误
	ErrorDatabase = 40001

	//对象存储错误
	ErrorOss        = 50001
	ErrorUploadFile = 50002
)
