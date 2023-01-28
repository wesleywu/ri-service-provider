package dubbogo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/WesleyWu/gf-cache/cache"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gtag"
	"reflect"
)

func init() {
	extension.SetFilter("InputValidationFilter", InputValidationFilter)
	extension.SetFilter("CacheFilter", CacheFilter)
}

func InputValidationFilter() filter.Filter {
	return &MyInputValidationFilter{}
}

func CacheFilter() filter.Filter {
	return &MyCacheFilter{}
}

type MyInputValidationFilter struct {
}

func (f *MyInputValidationFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	g.Log().Infof(ctx, "MyInputValidationFilter Invoke is called, method Name = %s", invocation.MethodName())
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
	err := result.Error()
	if err != nil {
		g.Log().Infof(ctx, "MyInputValidationFilter OnResponse error: %v", result.Error())
	}
	return result
}

type MyCacheFilter struct {
}

var ResTypes = map[string]reflect.Type{}

func mergeDefaultStructValue(pointer interface{}) error {
	// todo cache the reflect result
	defaultValueTags := []string{gtag.DefaultShort, gtag.Default}
	tagFields, err := gstructs.TagFields(pointer, defaultValueTags)
	if err != nil {
		return err
	}
	if len(tagFields) > 0 {
		//var (
		//	foundKey   string
		//	foundValue interface{}
		//)
		for _, field := range tagFields {
			if field.Value.IsZero() {
				field.Value.Set(reflect.ValueOf(gconv.Uint32(field.TagValue)))
			}
			//fieldValue := .Interface()
			//foundKey, foundValue = gutil.MapPossibleItemByKey(data, field.Name())
			//if foundKey == "" {
			//	data[field.Name()] = field.TagValue
			//} else {
			//	if empty.IsEmpty(foundValue) {
			//		data[foundKey] = field.TagValue
			//	}
			//}
		}
	}
	return nil
}

func (f *MyCacheFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	params := invocation.Arguments()
	methodName := invocation.MethodName()
	g.Log().Infof(ctx, "MyCacheFilter Invoke is called, method Name = %s, param value = %s", methodName, gjson.MustEncodeString(params[0]))
	switch methodName {
	case "List":
		// todo check if params len = 1
		req := params[0]
		reqType := reflect.TypeOf(params[0])
		// todo check reqType is a pointer to a struct
		// todo check if the struct has a field named 'Meta'
		// todo find cache settings in tag of field 'Meta'
		serviceName, exists := invocation.GetAttachment("interface")
		_ = mergeDefaultStructValue(req)
		g.Log().Infof(ctx, "param[0] type:%v, value:%v", reqType, params[0])
		g.Log().Infof(ctx, "serviceName:%v, exists:%v", serviceName, exists)
		if req == nil {
			break
		}
		if !cache.Initialized() {
			break
		}
		if resType, ok := ResTypes[serviceName+"_"+methodName]; ok {
			result := reflect.New(resType.Elem()).Interface()
			cacheKey := cache.GetCacheKey(serviceName, methodName, req)
			if cacheKey == nil {
				break
			}
			err := cache.RetrieveCacheTo(ctx, cacheKey, result)
			if err != nil {
				if err == cache.ErrLockTimeout { // 获取锁超时，返回降级的结果
					return &protocol.RPCResult{
						// todo add an attribute noting the result is a downgraded result
						Attrs: invocation.Attachments(),
						Err:   nil,
						Rest:  result,
					}
				} else if err == cache.ErrNotFound { // cache 未找到，执行底层操作
					break
				}
				// 其他底层错误
				g.Log().Error(ctx, err)
				break
			}
			// 返回缓存的结果
			return &protocol.RPCResult{
				// todo add an attribute noting the result is cached result
				Attrs: invocation.Attachments(),
				Err:   nil,
				Rest:  result,
			}
		}
	default:
		// todo implements other methods
	}
	return invoker.Invoke(ctx, invocation)
}
func (f *MyCacheFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	err := result.Error()
	if err != nil {
		g.Log().Infof(ctx, "MyInputValidationFilter OnResponse error: %v", result.Error())
	}
	params := protocol.Arguments()
	methodName := protocol.MethodName()
	serviceName, _ := protocol.GetAttachment("interface")
	cacheKey := cache.GetCacheKey(serviceName, methodName, params[0])
	if err == nil && cacheKey != nil && result != nil && cache.Initialized() {
		// todo if key exists, do not overwrite it
		ResTypes[serviceName+"_"+methodName] = reflect.TypeOf(result.Result())
		// if a cached result, no need to save back to cache again
		_ = cache.SaveCache(ctx, serviceName, cacheKey, result.Result())
	}
	return result
}
