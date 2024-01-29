/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
