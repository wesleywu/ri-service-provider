package gwerror

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type RequestError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	ReqBody string `json:"reqBody,omitempty"`
}

func (e RequestError) Error() string {
	return e.Message
}

func NewRequestErrorf(code int, req interface{}, format string, v ...any) RequestError {
	var (
		reqBody string
		ok      bool
	)
	reqBody = ""
	if req != nil {
		if reqBody, ok = req.(string); !ok {
			reqBody = gjson.MustEncodeString(req)
		}
	}
	return RequestError{
		Code:    code,
		Message: fmt.Sprintf(format, v...),
		ReqBody: reqBody,
	}
}

func NewBadRequestErrorf(req interface{}, format string, v ...any) RequestError {
	return NewRequestErrorf(400, req, format, v...)
}

func NewNotFoundErrorf(req interface{}, format string, v ...any) RequestError {
	return NewRequestErrorf(404, req, format, v...)
}

func NewPkConflictErrorf(req interface{}, format string, v ...any) RequestError {
	return NewRequestErrorf(409, req, format, v...)
}

func NewDataTooLongErrorf(req interface{}, format string, v ...any) RequestError {
	return NewRequestErrorf(413, req, format, v...)
}
