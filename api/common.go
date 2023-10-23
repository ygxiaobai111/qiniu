package api

import (
	"encoding/json"
	"www.github.com/ygxiaobai111/qiniu/serializer"
)

// ErrorResponse 定义返回错误格式
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}

}
