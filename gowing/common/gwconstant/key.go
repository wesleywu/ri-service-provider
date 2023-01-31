package gwconstant

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"strings"
)

const (
	CacheFilterKey           = "cache"
	SyncTraceFilterKey       = "sync-trace"
	OtelClientTraceFilterKey = "otel-client-trace"
	ValidationFilterKey      = "validation"
)

var (
	ServerFilters = strings.Join([]string{
		SyncTraceFilterKey,
		constant.EchoFilterKey,
		constant.MetricsFilterKey,
		constant.TokenFilterKey,
		constant.AccessLogFilterKey,
		constant.TpsLimitFilterKey,
		constant.GracefulShutdownProviderFilterKey,
		constant.TracingFilterKey,
		constant.OTELServerTraceKey,
		ValidationFilterKey,
		CacheFilterKey,
	}, ",")
	ClientFilters = strings.Join([]string{
		constant.GracefulShutdownConsumerFilterKey,
		OtelClientTraceFilterKey,
	}, ",")
)
