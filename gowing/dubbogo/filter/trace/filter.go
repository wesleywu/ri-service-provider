package trace

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/WesleyWu/ri-service-provider/gowing/common/gwconstant"
	"github.com/WesleyWu/ri-service-provider/gowing/util/gwmap"
	"github.com/gogf/gf/v2/net/gtrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	extension.SetFilter(gwconstant.SyncTraceFilterKey, func() filter.Filter {
		return &serverTraceFilter{}
	})
}

type serverTraceFilter struct {
}

func (f *serverTraceFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	traceId := gtrace.GetTraceID(ctx)
	if traceId == "" {
		attachments := invocation.Attachments()
		ctx = otel.GetTextMapPropagator().Extract(ctx, &gwmap.StrStrMap{
			InnerMap: attachments,
		})
		spanCtx := trace.SpanContextFromContext(ctx)
		traceId = spanCtx.TraceID().String()
		if traceId != "" {
			ctx, _ = gtrace.WithTraceID(ctx, traceId)
		}
	}
	return invoker.Invoke(ctx, invocation)
}
func (f *serverTraceFilter) OnResponse(_ context.Context, result protocol.Result, _ protocol.Invoker, _ protocol.Invocation) protocol.Result {
	return result
}
