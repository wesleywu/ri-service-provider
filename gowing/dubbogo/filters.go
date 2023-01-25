package dubbogo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	extension.SetFilter("InputValidationFilter", InputValidationFilter)
}

func InputValidationFilter() filter.Filter {
	return &MyInputValidationFilter{}
}

type MyInputValidationFilter struct {
}

func (f *MyInputValidationFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	fmt.Println("MyInputValidationFilter Invoke is called, method Name = ", invocation.MethodName())
	params := invocation.Arguments()
	for i, param := range params {
		err := g.Validator().Data(param).Run(ctx)
		if err != nil {
			return &protocol.RPCResult{
				Attrs: nil,
				Err:   err,
				Rest:  err.Error(),
			}
		}
		g.Log().Infof(ctx, "param %d: %v", i, param)
	}
	return invoker.Invoke(ctx, invocation)
}
func (f *MyInputValidationFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	fmt.Printf("MyInputValidationFilter OnResponse is called: %v\n", result.Error())
	return result
}
