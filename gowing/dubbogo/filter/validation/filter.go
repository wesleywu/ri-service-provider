package validation

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func init() {
	extension.SetFilter("validation", newFilter)
}

type validationFilter struct {
}

func newFilter() filter.Filter {
	return &validationFilter{}
}

func validateRequest(ctx context.Context, req interface{}) error {
	ctx, span := gtrace.NewSpan(ctx, "validateRequest")
	defer span.End()
	return g.Validator().Data(req).Run(ctx)
}

func (f *validationFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	ctx, span := gtrace.NewSpan(ctx, "validationFilter.Invoke")
	defer span.End()
	params := invocation.Arguments()
	for _, param := range params {
		err := validateRequest(ctx, param)
		if err != nil {
			return &protocol.RPCResult{
				Attrs: nil,
				Err:   err,
				Rest:  err.Error(),
			}
		}
		//g.Log().Infof(ctx, "param %d: %v", i, param)
	}
	return invoker.Invoke(ctx, invocation)
}
func (f *validationFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	ctx, span := gtrace.NewSpan(ctx, "validationFilter.OnResponse")
	defer span.End()
	err := result.Error()
	if err != nil {
		g.Log().Infof(ctx, "validationFilter OnResponse error: %v", result.Error())
	}
	return result
}
