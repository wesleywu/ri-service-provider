package gwhttpclient

import "net/http"

type HttpRequest struct {
	Method     string         // GET|POST|PUT|DELETE|OPTION
	Url        string         // 要请求的完整URL地址（包括协议、主机名、端口、Path路径、Query查询）
	Headers    *http.Header   // 请求使用的 Headers
	Cookies    []*http.Cookie // 请求使用的 Cookies
	Body       string         // 请求体字符串（GET|OPTION|DELETE请求传空字符串）
	RetryCount int            // 在成功之前，需要反复重试多少次，-1为重复1000次
}

type HttpResponse struct {
	Body       string       // 结果的Body字符串
	StatusCode int          // 结果的Http状态码
	Header     *http.Header // 结果的所有Http Header
}
