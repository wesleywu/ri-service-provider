package errors

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ServiceError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	ReqBody string `json:"reqBody,omitempty"`
	Err     error  `json:"err"`
}

type RequestError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	ReqBody string `json:"reqBody,omitempty"`
}

func (e ServiceError) Error() string {
	return e.Message + e.Err.Error()
}

func (e RequestError) Error() string {
	return e.Message
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
	return ServiceError{
		Code:    500,
		Message: fmt.Sprintf(format, v...),
		ReqBody: reqBody,
		Err:     err,
	}
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

func DbErrorToRequestError(req interface{}, err error, dbType string) (error, bool) {
	if err == nil {
		return nil, false
	}
	underlyingError := gerror.Unwrap(err)
	if underlyingError == nil {
		return nil, false
	}
	switch dbType {
	case "mysql":
		if driverError, ok := underlyingError.(*mysql.MySQLError); ok {
			switch driverError.Number {
			case 1062:
				return NewPkConflictErrorf(req, "%v: %v", "主键冲突", driverError.Error()), true
			case 1406:
				return NewDataTooLongErrorf(req, "%v: %v", "数据过长", driverError.Error()), true
			}
		}
	}
	return err, false
}
