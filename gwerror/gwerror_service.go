package gwerror

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type ServiceError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	ReqBody string `json:"reqBody,omitempty"`
	Err     error  `json:"err"`
}

func (e ServiceError) Error() string {
	return e.Message + e.Err.Error()
}
func WrapServiceErrorf(err error, req interface{}, format string, v ...any) ServiceError {
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
	code := 500
	if rerr, ok := err.(RequestError); ok {
		code = rerr.Code
	}
	return ServiceError{
		Code:    code,
		Message: fmt.Sprintf(format, v...),
		ReqBody: reqBody,
		Err:     err,
	}
}
