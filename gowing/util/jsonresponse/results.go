package jsonresponse

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	SuccessCode int = 0
	ErrorCode   int = -1
)

type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 错误消息
	Error string `json:"error,omitempty"`
	// 数据集
	Data interface{} `json:"data"`
}

func Success(r *ghttp.Request, data interface{}) {
	Result(r, SuccessCode, "", data)
}

func SuccessWithMessage(r *ghttp.Request, errorMessage string, data interface{}) {
	Result(r, SuccessCode, errorMessage, data)
}

func Failed(r *ghttp.Request, message string) {
	Result(r, ErrorCode, message, nil)
}

func FailedWithCode(r *ghttp.Request, code int, errorMessage string) {
	Result(r, code, errorMessage, nil)
}

func Result(r *ghttp.Request, code int, errorMessage string, data interface{}) {
	response := &Response{
		Code:  code,
		Error: errorMessage,
		Data:  data,
	}
	r.SetParam("apiReturnRes", response)
	r.Response.WriteJson(response)
	r.Exit()
}
