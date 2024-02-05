package benchmark

import (
	"net/http"
	"net/url"
	"testing"
)

func BenchmarkApiOne(b *testing.B) {
	httpClient := http.DefaultClient
	for i := 0; i < b.N; i++ {
		_, _ = httpClient.Get("http://localhost:8200/app/video-collection/one?id=01186883-7690")
	}
}

func BenchmarkApiList(b *testing.B) {
	httpClient := http.DefaultClient
	condition := "condition{\"operator\":\"Like\", \"wildcard\":\"StartsWith\", \"value\":\"推荐\"}"
	condition = url.QueryEscape(condition)
	for i := 0; i < b.N; i++ {
		_, _ = httpClient.Get("http://localhost:8200/app/video-collection/list?name=" + condition)
	}
}
