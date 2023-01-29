package cache

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/WesleyWu/ri-service-provider/gowing/cache"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gtag"
	"google.golang.org/protobuf/proto"
	"reflect"
)

const (
	CachedResult     = "cached-result"
	DowngradedResult = "downgraded-result"
)

func init() {
	extension.SetFilter("cache", newFilter)
}

func newFilter() filter.Filter {
	return &cacheFilter{}
}

type cacheFilter struct {
}

var ResTypes = map[string]reflect.Type{}

func mergeDefaultStructValue(ctx context.Context, pointer interface{}) error {
	ctx, span := gtrace.NewSpan(ctx, "mergeDefaultStructValue")
	defer span.End()
	// todo cache the reflect result
	defaultValueTags := []string{gtag.DefaultShort, gtag.Default}
	tagFields, err := gstructs.TagFields(pointer, defaultValueTags)
	if err != nil {
		return err
	}
	if len(tagFields) > 0 {
		for _, field := range tagFields {
			if field.Value.IsZero() {
				// todo switch tagValue's real type (not only Uint32)
				field.Value.Set(reflect.ValueOf(gconv.Uint32(field.TagValue)))
			}
		}
	}
	return nil
}

func (f *cacheFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	ctx, span := gtrace.NewSpan(ctx, "cacheFilter.Invoke")
	defer span.End()
	params := invocation.Arguments()
	service := invoker.GetURL().ServiceKey()
	methodName := invocation.ActualMethodName()
	g.Log().Infof(ctx, "cacheFilter Invoke is called, service name: %s, method Name: %s, param value: %s", service, methodName, gjson.MustEncodeString(params[0]))
	switch methodName {
	case "List":
		// todo check if params len = 1
		req := params[0].(proto.Message)
		//reqType := reflect.TypeOf(params[0])
		// todo check reqType is a pointer of protoMessage
		// todo check if the struct has a field named 'Meta'
		// todo find cache settings in tag of field 'Meta'
		_ = mergeDefaultStructValue(ctx, req)
		//g.Log().Infof(ctx, "param[0] type:%v, value:%v", reqType, params[0])
		//g.Log().Infof(ctx, "serviceName:%v, exists:%v", serviceName, exists)
		if req == nil {
			break
		}
		if !cache.Initialized() {
			break
		}
		if resType, ok := ResTypes[service+"_"+methodName]; ok {
			result := reflect.New(resType.Elem()).Interface().(proto.Message)
			cacheKey := cache.GetCacheKey(ctx, service, methodName, req)
			if cacheKey == nil {
				break
			}
			err := cache.RetrieveCacheTo(ctx, cacheKey, result)
			if err != nil {
				if err == cache.ErrLockTimeout { // 获取锁超时，返回降级的结果
					invocation.SetAttachment(DowngradedResult, true)
					return &protocol.RPCResult{
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
			invocation.SetAttachment(CachedResult, true)
			return &protocol.RPCResult{
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
func (f *cacheFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	ctx, span := gtrace.NewSpan(ctx, "cacheFilter.OnResponse")
	defer span.End()
	err := result.Error()
	if err != nil {
		g.Log().Infof(ctx, "CacheResponse OnResponse error: %v", result.Error())
	}
	if gconv.Bool(result.Attachment(CachedResult, false)) {
		return result
	}
	params := protocol.Arguments()
	service := invoker.GetURL().ServiceKey()
	methodName := protocol.ActualMethodName()
	cacheKey := cache.GetCacheKey(ctx, service, methodName, params[0].(proto.Message))
	if err == nil && cacheKey != nil && result.Result() != nil && cache.Initialized() {
		if _, exists := ResTypes[service+"_"+methodName]; !exists {
			ResTypes[service+"_"+methodName] = reflect.TypeOf(result.Result())
		}
		// if a cached result, no need to save back to cache again
		_ = cache.SaveCache(ctx, service, cacheKey, result.Result())
	}
	return result
}
