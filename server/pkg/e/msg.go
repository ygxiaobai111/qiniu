package e

// MsgFlags 错误详细信息
var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistInentifier:    "该第三方账号未绑定",
	ERRORNULL:             "无数据",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",

	ErrorExistNick:          "已存在该昵称",
	ErrorExistUser:          "已存在该用户名",
	ErrorNotExistUser:       "该用户不存在",
	ErrorNotCompare:         "账号密码错误",
	ErrorNotComparePassword: "两次密码输入不一致",
	ErrorFailEncryption:     "加密失败",

	ErrorUserNotFound: "用户不存在",

	ErrorUploadFile:    "上传失败",
	ErrorAdminFindUser: "管理员查询用户失败",

	ErrorDatabase: "数据库操作出错,请重试",

	ErrorOss: "OSS配置错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
