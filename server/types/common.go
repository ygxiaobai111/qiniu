package types

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	e2 "www.github.com/ygxiaobai111/qiniu/server/pkg/e"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type Page struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

// RespSuccess 带data成功返回
func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {

	status := e2.SUCCESS
	if code != nil {
		status = code[0]
	}

	if data == nil {
		data = "操作成功"
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e2.GetMsg(status),
	}

	return r
}

// ErrorResponse 定义返回错误格式
func ErrorResponse(err error) Response {
	if err == gorm.ErrRecordNotFound {
		return Response{
			Status: 404,
			Msg:    "并没有你想要的数据捏",
			Error:  err.Error(),
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}

}
