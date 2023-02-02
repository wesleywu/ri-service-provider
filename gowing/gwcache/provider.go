package gwcache

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/protobuf/proto"
)

const (
	ProviderNameMemory       = "memory"
	ProviderNameRedis        = "redis"
	ServiceCachePrefix       = "_SC_"
	ServiceCacheKeySetPrefix = "_SC_SET_"
	ServiceCacheLockerPrefix = "_LOCK_"
)

var (
	ctx                    = gctx.New()
	CacheEnabled           = false
	DebugEnabled           = false
	TracingEnabled         = false
	CacheProviderName      = "unknown"
	cacheProviderMap       = gmap.StrAnyMap{}
	ErrCacheNotInitialized = errors.New("cache: not initialized")
	ErrEmptyCacheKey       = errors.New("cache: cache key is empty")
	ErrLockTimeout         = errors.New("cache: lock timeout")
	ErrNotFound            = errors.New("cache: not found")
	ErrEmptyCachedValue    = errors.New("cache: cached value is empty")
)

type CacheProvider interface {
	Initialized() bool
	RetrieveCacheTo(ctx context.Context, cacheKey *string, value proto.Message) error
	SaveCache(ctx context.Context, serviceName string, cacheKey *string, value any, ttlSeconds uint32) error
	RemoveCache(ctx context.Context, cacheKey *string) error
	ClearCache(ctx context.Context, serviceName string) error
}

func init() {
	cacheEnabledVar, err := g.Cfg().Get(ctx, "cache.enabled", false)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return
	}
	cacheProviderVar, err := g.Cfg().Get(ctx, "cache.provider", "memory")
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return
	}
	debugEnabledVar, err := g.Cfg().Get(ctx, "cache.debug", false)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return
	}
	tracingEnabledVar, err := g.Cfg().Get(ctx, "cache.tracing", false)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return
	}
	CacheEnabled = cacheEnabledVar.Bool()
	DebugEnabled = debugEnabledVar.Bool()
	TracingEnabled = tracingEnabledVar.Bool()
	CacheProviderName = cacheProviderVar.String()
	if CacheEnabled {
		g.Log().Infof(ctx, "service cache enabled with provider '%s'", CacheProviderName)
	}
}

func createProviderIfNotExists(ctx context.Context, adapterType string, createProviderFunc func(context.Context) (CacheProvider, error)) (CacheProvider, error) {
	providerVar := cacheProviderMap.GetVarOrSetFuncLock(adapterType, func() interface{} {
		provider, err := createProviderFunc(ctx)
		if err != nil {
			return err
		}
		return provider
	})
	switch providerVar.Val().(type) {
	case error:
		return nil, providerVar.Val().(error)
	default:
		return providerVar.Val().(CacheProvider), nil
	}
}

func GetCacheProvider() (CacheProvider, error) {
	switch CacheProviderName {
	case ProviderNameMemory:
		return createProviderIfNotExists(ctx, ProviderNameMemory, NewMemoryCacheProvider)
	case ProviderNameRedis:
		return createProviderIfNotExists(ctx, ProviderNameRedis, NewRedisCacheProvider)
	default:
		return nil, gerror.Newf("Not supported cache provider: %s", CacheProviderName)
	}
}
