package cache

import (
	"context"
	"errors"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

type Storage interface {
	Initialized() bool
	Get(ctx context.Context, key string) ([]byte, error)
	GetString(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, content []byte, duration time.Duration) error
	SetString(ctx context.Context, key string, content string, duration time.Duration) error
	Delete(ctx context.Context, keys []string) error
	SAdd(ctx context.Context, setKey, value string) error
	SMembers(ctx context.Context, setKey string) ([]string, error)
}

var (
	RedisClient            *redis.Client
	RedisLocker            *redislock.Client
	LockTimeout            time.Duration
	CacheItemTtl           time.Duration
	storage                Storage
	ErrNilResult           = redis.Nil
	ErrCacheNotInitialized = errors.New("cache: not initialized")
	ErrEmptyCacheKey       = errors.New("cache: cache key is empty")
	ErrLockTimeout         = errors.New("cache: lock timeout")
	ErrNotFound            = errors.New("cache: not found")
)

func init() {
	ctx := gctx.New()
	lockTimeoutSeconds := g.Cfg().MustGet(ctx, "redis.default.lockTimeoutSeconds", 3).Int()
	LockTimeout = time.Duration(lockTimeoutSeconds) * time.Second
	cacheItemTtlMinutes := g.Cfg().MustGet(ctx, "redis.default.cacheItemTtlMinutes", 10).Int()
	CacheItemTtl = time.Duration(cacheItemTtlMinutes) * time.Minute
	address := g.Cfg().MustGet(ctx, "redis.default.host", "127.0.0.1:6379").String()
	db := g.Cfg().MustGet(ctx, "redis.default.db", 0).Int()
	password := g.Cfg().MustGet(ctx, "redis.default.password", "").String()
	RedisClient = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     address,
		DB:       db,
		Password: password,
	})
	info := RedisClient.Info(ctx)
	cacheInitialized := false
	if info.Err() != nil {
		g.Log().Warningf(ctx, "Error initialize redis: %s", info.Err().Error())
	} else {
		cacheInitialized = true
	}
	if RedisClient != nil {
		RedisLocker = redislock.New(RedisClient)
	}
	storage = &redisStorage{
		initialized: cacheInitialized,
		Client:      RedisClient,
	}
}

type redisStorage struct {
	initialized bool
	Client      *redis.Client
}

func (s *redisStorage) Initialized() bool {
	return s.initialized
}

func (s *redisStorage) Get(ctx context.Context, key string) ([]byte, error) {
	return s.Client.Get(ctx, key).Bytes()
}

func (s *redisStorage) GetString(ctx context.Context, key string) (string, error) {
	return s.Client.Get(ctx, key).Result()
}

func (s *redisStorage) Set(ctx context.Context, key string, content []byte, duration time.Duration) error {
	return s.Client.Set(ctx, key, content, duration).Err()
}

func (s *redisStorage) SetString(ctx context.Context, key string, content string, duration time.Duration) error {
	return s.Client.Set(ctx, key, content, duration).Err()
}

func (s *redisStorage) Delete(ctx context.Context, keys []string) error {
	return s.Client.Del(ctx, keys...).Err()
}

func (s *redisStorage) SAdd(ctx context.Context, setKey, value string) error {
	return s.Client.SAdd(ctx, setKey, value).Err()
}

func (s *redisStorage) SMembers(ctx context.Context, setKey string) ([]string, error) {
	return s.Client.SMembers(ctx, setKey).Result()
}
